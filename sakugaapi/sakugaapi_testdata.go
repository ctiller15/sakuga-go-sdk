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

var artistDataListResponse = []byte(`
	[
		{
			"id": 1172,
			"name": "심땅",
			"alias_id": 1171,
			"group_id": null,
			"urls": []
		},
		{
			"id": 1166,
			"name": "장석민",
			"alias_id": 1165,
			"group_id": null,
			"urls": []
		},
		{
			"id": 1170,
			"name": "파랑상자",
			"alias_id": 1169,
			"group_id": null,
			"urls": []
		},
		{
			"id": 647,
			"name": ".",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 672,
			"name": ">",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 593,
			"name": "0nepeach",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 1053,
			"name": "666ban",
			"alias_id": 1052,
			"group_id": null,
			"urls": []
		},
		{
			"id": 1192,
			"name": "abegen",
			"alias_id": 1191,
			"group_id": null,
			"urls": []
		},
		{
			"id": 1068,
			"name": "adam_zheng",
			"alias_id": 1066,
			"group_id": null,
			"urls": []
		},
		{
			"id": 651,
			"name": "aida_saakian",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 531,
			"name": "aimkid",
			"alias_id": null,
			"group_id": null,
			"urls": [
			"https://twitter.com/aimkidblast"
			]
		},
		{
			"id": 997,
			"name": "aimman_ibrahim",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 1027,
			"name": "aiowaruru",
			"alias_id": 616,
			"group_id": null,
			"urls": []
		},
		{
			"id": 799,
			"name": "aito_ohashi",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 1189,
			"name": "akapape",
			"alias_id": 811,
			"group_id": null,
			"urls": []
		},
		{
			"id": 782,
			"name": "akari_ranzaki",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 871,
			"name": "aki_deguchi",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 559,
			"name": "akihiro_ota",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 970,
			"name": "akiko_kudo",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 505,
			"name": "akio_watanabe",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 811,
			"name": "akira_hamaguchi",
			"alias_id": null,
			"group_id": null,
			"urls": [
			"https://x.com/oyasirazumonai",
			"https://x.com/akapape"
			]
		},
		{
			"id": 753,
			"name": "akira_kikuchi",
			"alias_id": null,
			"group_id": null,
			"urls": [
			"https://www.animenewsnetwork.com/encyclopedia/people.php?id=19539",
			"https://w.atwiki.jp/sakuga/pages/1138.amp",
			"https://w.atwiki.jp/anime_wiki/pages/6287.amp"
			]
		},
		{
			"id": 136,
			"name": "akira_takata",
			"alias_id": null,
			"group_id": null,
			"urls": []
		},
		{
			"id": 911,
			"name": "akira_tanaka",
			"alias_id": 910,
			"group_id": null,
			"urls": []
		},
		{
			"id": 975,
			"name": "ako_no_akio",
			"alias_id": null,
			"group_id": null,
			"urls": []
		}
	]
`)

var commentDataShowResponse = []byte(`
{
  "id": 90003,
  "created_at": "2023-12-26T20:52:33.419Z",
  "post_id": 107257,
  "creator": "Ivorybacon",
  "creator_id": 12904,
  "body": "I'm tempted to say that Otsuka is here given some of the smears on dispay. 0:05 and 0:11 for instance appear to have those very jagged smear lines that she seems to do often. (post #217312) and (post #215778) for comparison.\r\n\r\nShe's listed as an AD on this episode so there's a chance that it might just be a correction on her part."
}
`)
