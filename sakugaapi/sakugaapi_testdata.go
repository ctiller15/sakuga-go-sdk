package sakugaapi

var postDataResponse = []byte(`
[
  {
    "id": 216284,
    "tags": "artist_unknown correction kubo-san_wa_mob_wo_yurusanai production_materials",
    "created_at": 1674042317,
    "updated_at": 1674050468,
    "creator_id": 20764,
    "approver_id": 4456,
    "author": "JDMManga",
    "change": 1065989,
    "source": "https://twitter.com/PINEJAM_info/status/1615627553791672321?s=20&t=rThR9qaXQaB6JtaaPczLoA",
    "score": 0,
    "md5": "94818e87adafcd4a8b857eca97417b69",
    "file_size": 2607352,
    "file_ext": "png",
    "file_url": "https://www.sakugabooru.com/data/94818e87adafcd4a8b857eca97417b69.png",
    "is_shown_in_index": true,
    "preview_url": "https://www.sakugabooru.com/data/preview/94818e87adafcd4a8b857eca97417b69.jpg",
    "preview_width": 150,
    "preview_height": 106,
    "actual_preview_width": 300,
    "actual_preview_height": 212,
    "sample_url": "https://www.sakugabooru.com/data/sample/94818e87adafcd4a8b857eca97417b69.jpg",
    "sample_width": 1500,
    "sample_height": 1061,
    "sample_file_size": 237252,
    "jpeg_url": "https://www.sakugabooru.com/data/94818e87adafcd4a8b857eca97417b69.png",
    "jpeg_width": 7484,
    "jpeg_height": 5292,
    "jpeg_file_size": 0,
    "rating": "s",
    "is_rating_locked": false,
    "has_children": false,
    "parent_id": null,
    "status": "active",
    "is_pending": false,
    "width": 7484,
    "height": 5292,
    "is_held": false,
    "frames_pending_string": "",
    "frames_pending": [],
    "frames_string": "",
    "frames": [],
    "is_note_locked": false,
    "last_noted_at": 0,
    "last_commented_at": 0
  }
]
`)

var tagDataResponse = []byte(`
[
  {
    "id": 14391,
    "name": "",
    "count": 0,
    "type": 0,
    "ambiguous": false
  },
  {
    "id": 25325,
    "name": "(???)",
    "count": 0,
    "type": 0,
    "ambiguous": false
  },
  {
    "id": 25384,
    "name": "*",
    "count": 0,
    "type": 0,
    "ambiguous": false
  },
  {
    "id": 22702,
    "name": "?",
    "count": 0,
    "type": 0,
    "ambiguous": false
  },
  {
    "id": 24506,
    "name": "\\",
    "count": 0,
    "type": 0,
    "ambiguous": false
  },
  {
    "id": 26578,
    "name": "]",
    "count": 0,
    "type": 0,
    "ambiguous": false
  },
  {
    "id": 25768,
    "name": "_",
    "count": 0,
    "type": 0,
    "ambiguous": false
  },
  {
    "id": 25547,
    "name": "0:00/0:46",
    "count": 0,
    "type": 0,
    "ambiguous": false
  },
  {
    "id": 5620,
    "name": "00:08",
    "count": 5,
    "type": 3,
    "ambiguous": false
  },
  {
    "id": 11330,
    "name": "001",
    "count": 1,
    "type": 3,
    "ambiguous": false
  }
]
`)

var tagDataRelatedResponse = []byte(
	`{
  "animated": [
    [
      "animated",
      "155193"
    ],
    [
      "effects",
      "88897"
    ],
    [
      "artist_unknown",
      "81238"
    ],
    [
      "character_acting",
      "60788"
    ],
    [
      "smoke",
      "42239"
    ],
    [
      "smears",
      "41842"
    ],
    [
      "fighting",
      "38855"
    ],
    [
      "creatures",
      "26170"
    ],
    [
      "presumed",
      "23921"
    ],
    [
      "liquid",
      "20841"
    ],
    [
      "debris",
      "18986"
    ],
    [
      "explosions",
      "16904"
    ],
    [
      "hair",
      "16480"
    ],
    [
      "running",
      "15126"
    ],
    [
      "fabric",
      "15021"
    ],
    [
      "western",
      "14087"
    ],
    [
      "background_animation",
      "12432"
    ],
    [
      "mecha",
      "12013"
    ],
    [
      "fire",
      "11356"
    ],
    [
      "sparks",
      "10068"
    ],
    [
      "animals",
      "9591"
    ],
    [
      "impact_frames",
      "9347"
    ],
    [
      "lightning",
      "9244"
    ],
    [
      "cgi",
      "8381"
    ],
    [
      "production_materials",
      "7688"
    ]
  ]
}`)
