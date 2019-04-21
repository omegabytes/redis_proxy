package cache

import (
	"fmt"
	"os"
	"redis_proxy/app/db"
	"redis_proxy/app/models"
	"strconv"
	"sync"
	"time"
)

var (
	SIZE, _ = strconv.Atoi(os.Getenv("CACHE_SIZE"))
	TTL, _  = time.ParseDuration(os.Getenv("CACHE_TTL") + "s")
)

type Node struct {
	sync.RWMutex
	Val     *models.Post
	Expires *time.Time
	Left    *Node
	Right   *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

// maps post item to node in Queue
type Hash map[string]*Node

type Cache struct {
	mutex sync.RWMutex
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(postId string) *models.Post {
	node := &Node{}
	c.mutex.Lock()

	if val, ok := c.Hash[postId]; ok {
		if val.Expired() {
			fmt.Println("[DEBUG] expired, adding to cache")
			post := db.GetPost(postId)

			if post != nil {
				node = &Node{Val: post}
				node.SetTTL()
			} else {
				c.mutex.Unlock()
				return nil
			}
		} else {
			fmt.Printf("[DEBUG] getting from cache (expires: %s)\n", *val.Expires)
			node = c.Remove(val)
		}
	} else {
		fmt.Println("[DEBUG] adding to cache")
		post := db.GetPost(postId)

		if post != nil {
			node = &Node{Val: post}
			node.SetTTL()
		} else {
			c.mutex.Unlock()
			return nil
		}
	}

	c.Add(node)
	c.Hash[postId] = node

	c.mutex.Unlock()
	return node.Val
}

func (c *Cache) Remove(n *Node) *Node {
	left := n.Left
	right := n.Right
	left.Right = right
	right.Left = left
	c.Queue.Length -= 1

	delete(c.Hash, n.Val.Title)
	return n
}

func (c *Cache) Add(n *Node) {
	tmp := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Display() {
	c.mutex.Lock()
	c.Queue.Display()
	c.mutex.Unlock()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("[DEBUG] len(%d) - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("%s", *node.Val)
		if i < q.Length-1 {
			fmt.Printf(", ")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func (n *Node) SetTTL() {
	n.Lock()
	expiration := time.Now().Add(TTL)
	n.Expires = &expiration
	n.Unlock()
}

func (n *Node) Expired() bool {
	var value bool
	n.RLock()
	if n.Expires == nil {
		value = true
	} else {
		value = n.Expires.Before(time.Now())
	}
	n.RUnlock()
	return value
}
