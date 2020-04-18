package main

import (
	"io"
)

// Node structure
type Node struct {    
	Type                       NodeType
	Data                       string
	Attr                       []Attribute
	FirstChild, NextSibling    *Node    
}

// NodeType equals to int32 
type NodeType int32

const (
	// ErrorNode Node type
	ErrorNode NodeType  = iota
	// TextNode ..
	TextNode
	// DocumentNode ..
	DocumentNode
	// ElementNode ..
	ElementNode
	// CommentNode ..
	CommentNode
	// DoctypeNode ..
	DoctypeNode

)

// Attribute structure
type Attribute struct {
	Key, Val string
}

// Parse Function
func Parse(r io.Reader) (*Node, error){

}

