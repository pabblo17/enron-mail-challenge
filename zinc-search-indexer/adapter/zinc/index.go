package zinc

func GetIndex() string {
	/*return `{
	"name": "mail",
	"storage_type": "disk",
	"shard_num": 1,
	"settings": {
		"analysis": {
			"analyzer": {
				"default": {
					"type": "standard"
				}
			}
		}
	},*/
	return `{
	"name": "mail",
	"storage_type": "disk",
	"shard_num": 1,
	"mappings": {
		"properties": {
			"message-id": {
				"type": "text",
				"index": false,
				"store": false
			},
			"from": {
				"type": "text",
				"index": false,
				"store": false
			},
			"to": {
				"type": "text",
				"index": false,
				"store": false
			},
			"subject": {
				"type": "text",
				"index": true,
				"store": false
			},
			"cc": {
				"type": "text",
				"index": false,
				"store": false
			},
			
			"bcc": {
				"type": "text",
				"index": false,
				"store": false
			},
			
			"X-from": {
				"type": "text",
				"index": false,
				"store": false
			},
			
			"x-to": {
				"type": "text",
				"index": false,
				"store": false
			},
			
			"x-cc": {
				"type": "text",
				"index": false,
				"store": false
			},
			"x-bcc": {
				"type": "text",
				"index": false,
				"store": false
			},
			"content": {
				"type": "text",
				"index": true,
				"store": false,
				"highlightable": true
			},
			
			"path": {
				"type": "text",
				"index": false,
				"store": false
			},
			"mode": {
				"type": "text",
				"index": false,
				"store": false
			}
		}
	}
}`

}
