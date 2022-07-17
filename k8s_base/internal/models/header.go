package models

type Header [2]string

var (
	ContentTypeJson = Header{"Content-Type", "json"}
	CharsetUft8     = Header{"charset", "uft-8"}
)
