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

var wikiDataListResponse = []byte(`
[
  {
    "id": 42,
    "created_at": "2016-10-05T13:30:31.967Z",
    "updated_at": "2016-10-05T13:30:31.967Z",
    "title": "3d_background",
    "body": "Background is made with 3dcg. Implicates [[cgi]].",
    "updater_id": 933,
    "locked": false,
    "version": 1
  },
  {
    "id": 312,
    "created_at": "2024-11-01T05:17:08.294Z",
    "updated_at": "2024-11-01T05:17:08.294Z",
    "title": "aimman_ibrahim",
    "body": "Brother of \"Jack-Amin Ibrahim\":https://www.sakugabooru.com/post?tags=jack-amin_ibrahim",
    "updater_id": 34147,
    "locked": false,
    "version": 1
  },
  {
    "id": 58,
    "created_at": "2017-08-30T11:10:14.286Z",
    "updated_at": "2017-08-30T11:43:21.687Z",
    "title": "ai_monogatari:_9_love_stories",
    "body": "h2. 「HERO」\r\n\r\n監督・キャラクターデザイン \r\n森本晃司 Kouji Morimoto\r\n\r\nキャラクターデザイン・作画監督\r\n清山滋崇 Shigetaka Kiyoyama\r\n\r\n原画\r\n誟橋伸司 Shinji Morohashi\r\n後藤孝宏 Takahiro Gotou\r\n西村　智 Satoshi Nishimura\r\n安藤正浩 Masahiro Andou (the other one)\r\n田辺　修 Osamu Tanabe\r\n結城明宏 Akihiro Yuuki\r\n\r\n\r\n\r\n\r\nh2. 「夜をぶっとばせ」\r\n\r\n監督\r\n浜津　守 Mamoru Hamatsu\r\n\r\nキャラクターデザイン・作画監督・原画\r\n村瀬修功 Shuukou Murase\r\n\r\n\r\n\r\n\r\nh2. 「ライオンとペリカン」\r\n\r\n監督\r\n澤井幸次 Kouji Sawai\r\n\r\nキャラクタデザイン・作画監督\r\n戸倉紀元 Norimoto Tokura\r\n\r\n原画\r\n戸倉紀元 Norimoto Tokura\r\n藤川　太 Futoshi Fujikawa\r\n\r\n",
    "updater_id": 933,
    "locked": false,
    "version": 4
  },
  {
    "id": 211,
    "created_at": "2022-11-12T12:51:55.235Z",
    "updated_at": "2022-11-12T12:51:55.235Z",
    "title": "aito_ohashi",
    "body": "大橋 藍人",
    "updater_id": 6259,
    "locked": false,
    "version": 1
  },
  {
    "id": 202,
    "created_at": "2022-10-03T14:31:32.798Z",
    "updated_at": "2022-10-03T14:31:32.798Z",
    "title": "akari_ranzaki",
    "body": "藍崎　灯",
    "updater_id": 6259,
    "locked": false,
    "version": 1
  }
]
`)

var noteDataListResponse = []byte(`
[
  {
    "id": 270,
    "created_at": "2024-02-06T22:15:57.304Z",
    "updated_at": "2024-02-06T22:15:57.304Z",
    "creator_id": 4456,
    "x": 555,
    "y": 137,
    "width": 153,
    "height": 119,
    "is_active": true,
    "post_id": 248418,
    "body": "object is always cel\nfull color trace paint",
    "version": 1
  },
  {
    "id": 269,
    "created_at": "2024-02-06T22:08:57.131Z",
    "updated_at": "2024-02-06T22:08:57.131Z",
    "creator_id": 4456,
    "x": 277,
    "y": 5,
    "width": 148,
    "height": 148,
    "is_active": true,
    "post_id": 248418,
    "body": "Demonic Mirroring Ice Crystals ",
    "version": 1
  }
]
`)

var noteDataSearchResponse = []byte(`
[
  {
    "id": 4,
    "created_at": "2013-09-03T20:29:32.200Z",
    "updated_at": "2013-09-03T20:29:32.200Z",
    "creator_id": 55,
    "x": 28,
    "y": 21,
    "width": 861,
    "height": 100,
    "is_active": true,
    "post_id": 1273,
    "body": "        \"Well, even if I do stuff like this, there probably isn't anyone who understands what it means.\"\n         That was the first time I'd heard him make that complaint since the \"Moonrise Mannor\" screening.\n        I imagine those complaints of his, as he cheerfully streaked through a public street in the middle of winter, were his true feelings.\n        The life of Ikuhara Kunihiko already had no distinction between public and private.\n \n        Just what is this guy, who dares to dance in the nude, like in the other world, where he's really going for it?\n        The things he fights against are instinct and normalcy. I think those two are the big ones.\n        Basically, concern for others is part of it. In most cases, people are unconsciously dependent on those sorts of things.\n        But, Ikuhara Kunihiko just seems to intuitively understand that using those elements as a tools is an inherent part of being human as well. I think he already had the \"subjective goal\" necessary to do that as a student. In other words, a spire of determination.\n \n        So––\n       \n        Now, he seems to have rid himself of his discontent in not finding anyone who understands him. But, in no way do I mean to imply that was the result of finding a friend he'd always hoped for. It's really quite the opposite. It's more that, he's become acutely cognizant of the fact that everyone other than himself is an enemy.\n        Banished from paradise, man was made mortal.\n        I think, Ikuhara Kunihiko decides to \"be Ikuhara Kunihiko\" in film, TV and even in his New Year's cards, with that harsh reality firmly planted in his mind.\n        He was already well aware of the struggle with the truth that you are you and nothing more.\n        But, I desperately wish that I might be able to watch over what he does from now on, firm in the belief that there is some truth to that stance of his.\n        Peace.\n \n12/21/1994\n \n---\nNotes:\n \n[1] - I did some digging and I believe this story pertains to episode 31 of Sailor Moon, \"恋されて追われて!ルナの最悪の日.\" It won the Animage Anime Grand Prix for individual episode of a TV series in 1992. Ikuhara storyboarded it and it does involve cats and mice (it's that Luna-centric episode).\n \nReferences:\nhttp://ja.wikipedia.org/wiki/アニメグランプリ\nhttp://ja.wikipedia.org/wiki/美少女戦士セーラームーン_(アニメ)\n \n \n[2] - 製作進行: Apparently this Toei's unique nomenclature for 制作進行. \"Production Assistant\" or \"Production Runner\" seem to be the common English translations for the position. I've also seen \"Production Coordinator.\" According to Wikipedia, at least, the job is often a stepping stone to becoming a producer or director, so they do a variety of odd jobs to help out the main staff and keep things running smoothly.\n \nReferences:\nhttp://ja.wikipedia.org/wiki/%E5%88%B6%E4%BD%9C%E9%80%B2%E8%A1%8C#.E3.82.A2.E3.83.8B.E3.83.A1.E3.83.BC.E3.82.B7.E3.83.A7.E3.83.B3.E4.BD.9C.E5.93.81.E3.81.AE.E5.88.B6.E4.BD.9C.E9.80.B2.E8.A1.8C\n \n \n[3] - The last card in the set they copied for this book.",
    "version": 1
  },
  {
    "id": 70,
    "created_at": "2016-12-12T00:01:36.576Z",
    "updated_at": "2016-12-12T00:01:36.576Z",
    "creator_id": 933,
    "x": 250,
    "y": 180,
    "width": 90,
    "height": 150,
    "is_active": true,
    "post_id": 28259,
    "body": "Since then he's done the scene in Gall Force where monsters come out of the crumbling ship, and the scene at the end of Kikoukai Galient: Tetsu no Monshou where Galient cuts Jashinhei apart. In Machine Robo, he did Baikanfuu's henkei which became a bank cut, Baikanfuu and Mizuchi's fight scene in episode 6 etc.",
    "version": 1
  },
  {
    "id": 71,
    "created_at": "2016-12-12T00:08:26.418Z",
    "updated_at": "2016-12-12T00:08:26.418Z",
    "creator_id": 933,
    "x": 202,
    "y": 176,
    "width": 47,
    "height": 159,
    "is_active": true,
    "post_id": 28259,
    "body": "Recently he's done the fight scene with the monsters in Gakuen Tokusou Hikaruon and the baseball scene in Bats and Terry, and been involved in Bubblegum Crisis, among others.",
    "version": 1
  }
]
`)

var userDataSearchResponse = []byte(`
[
  {
    "name": "sakurga56",
    "id": 39722
  },
  {
    "name": "Kanegane",
    "id": 39721
  },
  {
    "name": "bonofyah",
    "id": 39720
  },
  {
    "name": "antoines_workshop",
    "id": 39719
  },
  {
    "name": "Mekkies",
    "id": 39718
  },
  {
    "name": "Phoslos",
    "id": 39717
  },
  {
    "name": "gugull",
    "id": 39716
  },
  {
    "name": "ForestElephant",
    "id": 39715
  },
  {
    "name": "suibo",
    "id": 39714
  },
  {
    "name": "qJvECnDaBJcH",
    "id": 39713
  },
  {
    "name": "Etozy",
    "id": 39712
  },
  {
    "name": "奥利奥",
    "id": 39711
  },
  {
    "name": "alazoro",
    "id": 39710
  },
  {
    "name": "nikomi_shio",
    "id": 39709
  },
  {
    "name": "Barusu",
    "id": 39708
  },
  {
    "name": "GEEanimations",
    "id": 39707
  },
  {
    "name": "GT",
    "id": 39706
  },
  {
    "name": "MARL_",
    "id": 39705
  },
  {
    "name": "Aley",
    "id": 39704
  },
  {
    "name": "anzu",
    "id": 39703
  }
]
`)

var forumDataListResponse = []byte(`
  {
    "body": "If the website breaks or does something unexpected, post about it here.\r\n\r\nYou can also contact the booru Admin team via the following channels:\r\n\r\naers - http://sakuga.yshi.org/user/show/1 / https://twitter.com/aers00\r\nkvin - http://sakuga.yshi.org/user/show/2 / https://twitter.com/Yuyucow\r\nme - http://sakuga.yshi.org/user/show/4 / https://twitter.com/Kraker2k\r\n\r\nYou can also visit #sakuga on irc.rizon.net and leave a message there.\r\n\r\n",
    "creator": "Kraker2k",
    "creator_id": 4,
    "id": 12,
    "parent_id": null,
    "title": "Post errors and admin questions here",
    "updated_at": "2024-08-29T21:13:33.244Z",
    "pages": 2
  },
    {
    "body": "is there a way to search by how long or how big it is?\r\nlike longer than/shorter than  seconds\r\nor by size bigger/smaller\r\nalso is there a way to order the anime by how many posts in their tags there is or something",
    "creator": "amirdrama",
    "creator_id": 24087,
    "id": 1580,
    "parent_id": null,
    "title": "is there a way to search by how long or how big it is?",
    "updated_at": "2024-04-14T19:32:01.866Z",
    "pages": 1
  },
  {
    "body": " Anyone know of a community or place to get feedback and talk about creating animation? Like a serious one with good activity, not full of people that look like they drew stuff in MS paint (no offense to them). Like a discord or website.",
    "creator": "dracobuster",
    "creator_id": 24948,
    "id": 1352,
    "parent_id": null,
    "title": "Practicing Animation",
    "updated_at": "2024-04-01T16:51:22.984Z",
    "pages": 1
  },
  {
    "body": "Hello, several months ago I created a pool named \"Regeneration\". Where is it now? Is it violated some rules or something like?",
    "creator": "Nojtovsky",
    "creator_id": 19519,
    "id": 1563,
    "parent_id": null,
    "title": "Where is my pool?",
    "updated_at": "2024-04-01T15:26:09.392Z",
    "pages": 1
  },
  {
    "body": "is there a place to archive anime official arts?  \r\nlike anime Blu-ray covers/arts that are for magazine etc...",
    "creator": "amirdrama",
    "creator_id": 24087,
    "id": 1473,
    "parent_id": null,
    "title": "is there a place to archive anime official arts?",
    "updated_at": "2024-03-22T14:44:51.107Z",
    "pages": 1
  },
  {
    "body": "I really liked this cut in AoT final season but CG police took it down for a very unjustified reason imo. I personally don't know how best to record/reupload this\r\nhttps://www.sakugabooru.com/post/show/241415",
    "creator": "ItsAMemeMario",
    "creator_id": 23275,
    "id": 1565,
    "parent_id": null,
    "title": "Any way to bring a deleted post back?",
    "updated_at": "2024-03-06T12:02:40.117Z",
    "pages": 1
  }
`)

var poolDataListResponse = []byte(`
[
  {
    "id": 257,
    "name": "PPURI",
    "created_at": "2024-10-19T17:58:29.325Z",
    "updated_at": "2025-01-05T01:31:19.717Z",
    "user_id": 1028,
    "is_public": true,
    "post_count": 300,
    "description": "All the animation works from the Korean studio."
  },
  {
    "id": 249,
    "name": "Stair-Climbing",
    "created_at": "2024-06-14T17:48:09.391Z",
    "updated_at": "2025-02-01T01:40:41.690Z",
    "user_id": 13069,
    "is_public": true,
    "post_count": 37,
    "description": "Specialized *character acting* featuring stair-climbing, is generally considered to be one of the most difficult types of character acting to draw well. Preferably contains challenging/interesting layouts and movements (of the character acting) that capitalizes on the stair setting. Would prefer to avoid sequences where the movement is handled similarly to a standard run cycle."
  },
  {
    "id": 246,
    "name": "Botanical_Sakuga",
    "created_at": "2024-04-05T16:50:37.446Z",
    "updated_at": "2025-02-20T15:07:19.790Z",
    "user_id": 18138,
    "is_public": true,
    "post_count": 111,
    "description": "Well animated scenes of plants!"
  },
  {
    "id": 237,
    "name": "Kyoani_playing_cards",
    "created_at": "2024-03-07T16:29:11.632Z",
    "updated_at": "2025-01-23T20:53:32.593Z",
    "user_id": 21916,
    "is_public": true,
    "post_count": 31,
    "description": "パラパラトランプ"
  },
  {
    "id": 232,
    "name": "Sign_Language",
    "created_at": "2024-01-02T11:50:25.013Z",
    "updated_at": "2024-11-19T22:51:11.781Z",
    "user_id": 20764,
    "is_public": true,
    "post_count": 47,
    "description": ""
  }
]
`)

var poolShowDataResponse = []byte(`
{
  "id": 215,
  "name": "Stick_Figure_Animations",
  "created_at": "2023-06-21T19:45:51.333Z",
  "updated_at": "2024-08-08T21:07:28.095Z",
  "user_id": 18138,
  "is_public": false,
  "post_count": 20,
  "description": "Animations containing Stick Figures. Whether it'd fighting, dancing or mucking about, in general.",
  "posts": [
    {
      "id": 22733,
      "tags": "animated fighting philips_lacanlale smears web western",
      "created_at": "2016-05-06T01:23:30.476Z",
      "creator_id": 995,
      "author": "kintuin",
      "change": 1156350,
      "source": "",
      "score": 124,
      "md5": "92da968affd0ab4a7a31e4fb7e2755f6",
      "file_size": 884515,
      "file_url": "https://www.sakugabooru.com/data/92da968affd0ab4a7a31e4fb7e2755f6.gif",
      "is_shown_in_index": true,
      "preview_url": "https://www.sakugabooru.com/data/preview/92da968affd0ab4a7a31e4fb7e2755f6.jpg",
      "preview_width": 150,
      "preview_height": 105,
      "actual_preview_width": 300,
      "actual_preview_height": 210,
      "sample_url": "https://www.sakugabooru.com/data/92da968affd0ab4a7a31e4fb7e2755f6.gif",
      "sample_width": 500,
      "sample_height": 350,
      "sample_file_size": 0,
      "jpeg_url": "https://www.sakugabooru.com/data/92da968affd0ab4a7a31e4fb7e2755f6.gif",
      "jpeg_width": 500,
      "jpeg_height": 350,
      "jpeg_file_size": 0,
      "rating": "s",
      "has_children": false,
      "parent_id": null,
      "status": "active",
      "width": 500,
      "height": 350,
      "is_held": false,
      "frames_pending_string": "",
      "frames_pending": [],
      "frames_string": "",
      "frames": []
    },
    {
      "id": 29421,
      "tags": "animated beams debris effects fighting smears smoke tommy web",
      "created_at": "2017-01-05T10:11:41.217Z",
      "creator_id": 508,
      "author": "Ashita",
      "change": 1161941,
      "source": "",
      "score": 42,
      "md5": "7c4ea5f75f6901e2ec5bb4c8289a9e93",
      "file_size": 1179880,
      "file_url": "https://www.sakugabooru.com/data/7c4ea5f75f6901e2ec5bb4c8289a9e93.mp4",
      "is_shown_in_index": true,
      "preview_url": "https://www.sakugabooru.com/data/preview/7c4ea5f75f6901e2ec5bb4c8289a9e93.jpg",
      "preview_width": 150,
      "preview_height": 113,
      "actual_preview_width": 256,
      "actual_preview_height": 192,
      "sample_url": "https://www.sakugabooru.com/data/7c4ea5f75f6901e2ec5bb4c8289a9e93.mp4",
      "sample_width": 256,
      "sample_height": 192,
      "sample_file_size": 0,
      "jpeg_url": "https://www.sakugabooru.com/data/7c4ea5f75f6901e2ec5bb4c8289a9e93.mp4",
      "jpeg_width": 256,
      "jpeg_height": 192,
      "jpeg_file_size": 0,
      "rating": "s",
      "has_children": false,
      "parent_id": null,
      "status": "active",
      "width": 256,
      "height": 192,
      "is_held": false,
      "frames_pending_string": "",
      "frames_pending": [],
      "frames_string": "",
      "frames": []
    },
    {
      "id": 135811,
      "tags": "animated animator_vs._animation_8:_the_showdown animator_vs._animation_series artist_unknown beams effects explosions fighting fire flying guzzu impact_frames liquid shadowqrow smears smoke web western",
      "created_at": "2020-10-26T02:32:47.778Z",
      "creator_id": 5084,
      "author": "Armando",
      "change": 1323223,
      "source": "https://www.youtube.com/watch?v=0a1r0JaONS4",
      "score": 115,
      "md5": "277c0d99ef7aa2c9a5c443aa7580f0d0",
      "file_size": 20868373,
      "file_url": "https://www.sakugabooru.com/data/277c0d99ef7aa2c9a5c443aa7580f0d0.mp4",
      "is_shown_in_index": true,
      "preview_url": "https://www.sakugabooru.com/data/preview/277c0d99ef7aa2c9a5c443aa7580f0d0.jpg",
      "preview_width": 150,
      "preview_height": 84,
      "actual_preview_width": 300,
      "actual_preview_height": 169,
      "sample_url": "https://www.sakugabooru.com/data/277c0d99ef7aa2c9a5c443aa7580f0d0.mp4",
      "sample_width": 854,
      "sample_height": 480,
      "sample_file_size": 0,
      "jpeg_url": "https://www.sakugabooru.com/data/277c0d99ef7aa2c9a5c443aa7580f0d0.mp4",
      "jpeg_width": 854,
      "jpeg_height": 480,
      "jpeg_file_size": 0,
      "rating": "s",
      "has_children": false,
      "parent_id": null,
      "status": "active",
      "width": 854,
      "height": 480,
      "is_held": false,
      "frames_pending_string": "",
      "frames_pending": [],
      "frames_string": "",
      "frames": []
    },
    {
      "id": 135820,
      "tags": "animated animator_vs._animation_8:_the_showdown animator_vs._animation_series artist_unknown beams debris effects explosions falling fighting fire flying guzzu liquid rotation shadowqrow smears smoke web western",
      "created_at": "2020-10-26T04:38:51.279Z",
      "creator_id": 5084,
      "author": "Armando",
      "change": 1323224,
      "source": "https://www.youtube.com/watch?v=0a1r0JaONS4",
      "score": 168,
      "md5": "f4078f9ef536d04bb6d3e038beb94547",
      "file_size": 19062220,
      "file_url": "https://www.sakugabooru.com/data/f4078f9ef536d04bb6d3e038beb94547.mp4",
      "is_shown_in_index": true,
      "preview_url": "https://www.sakugabooru.com/data/preview/f4078f9ef536d04bb6d3e038beb94547.jpg",
      "preview_width": 150,
      "preview_height": 84,
      "actual_preview_width": 300,
      "actual_preview_height": 169,
      "sample_url": "https://www.sakugabooru.com/data/f4078f9ef536d04bb6d3e038beb94547.mp4",
      "sample_width": 854,
      "sample_height": 480,
      "sample_file_size": 0,
      "jpeg_url": "https://www.sakugabooru.com/data/f4078f9ef536d04bb6d3e038beb94547.mp4",
      "jpeg_width": 854,
      "jpeg_height": 480,
      "jpeg_file_size": 0,
      "rating": "s",
      "has_children": true,
      "parent_id": null,
      "status": "active",
      "width": 854,
      "height": 480,
      "is_held": false,
      "frames_pending_string": "",
      "frames_pending": [],
      "frames_string": "",
      "frames": []
    },
    {
      "id": 135822,
      "tags": "animated animator_vs._animation_8:_the_showdown animator_vs._animation_series artist_unknown effects fighting fire flying guzzu lightning liquid shadowqrow smears smoke sparks web western wind",
      "created_at": "2020-10-26T04:41:05.045Z",
      "creator_id": 5084,
      "author": "Armando",
      "change": 1323225,
      "source": "https://www.youtube.com/watch?v=0a1r0JaONS4",
      "score": 58,
      "md5": "1f9314adcc612ede1a25692b37b63647",
      "file_size": 13446784,
      "file_url": "https://www.sakugabooru.com/data/1f9314adcc612ede1a25692b37b63647.mp4",
      "is_shown_in_index": true,
      "preview_url": "https://www.sakugabooru.com/data/preview/1f9314adcc612ede1a25692b37b63647.jpg",
      "preview_width": 150,
      "preview_height": 84,
      "actual_preview_width": 300,
      "actual_preview_height": 169,
      "sample_url": "https://www.sakugabooru.com/data/1f9314adcc612ede1a25692b37b63647.mp4",
      "sample_width": 854,
      "sample_height": 480,
      "sample_file_size": 0,
      "jpeg_url": "https://www.sakugabooru.com/data/1f9314adcc612ede1a25692b37b63647.mp4",
      "jpeg_width": 854,
      "jpeg_height": 480,
      "jpeg_file_size": 0,
      "rating": "s",
      "has_children": false,
      "parent_id": null,
      "status": "active",
      "width": 854,
      "height": 480,
      "is_held": false,
      "frames_pending_string": "",
      "frames_pending": [],
      "frames_string": "",
      "frames": []
    }
	]
}
`)

var favoriteUsersData = []byte(`
{
  "favorited_users": "Kawaki,Zzdapao,Nave,Jupiterjavelin,Moondoggie,Dominic_Falisi,Anitrix,Petrychrores,SandyWick,mg13,kaise2347,m2,Miyagi,Ravennnn,Spicygolf,Fy3,DanielSonic,Stouze,Mokuga,Clon0s,Nic2007,martingrasoso,Werdert,swimhigh,sakinreyy,lvl2frog,sakugayballs,Brubthevideogameduck,drake366,mzamecki,Miyuki,goatt,Welipe,mybigheartforyou,29dan5,Jen,akina432,poipop_tiktok,Martin_Caubet,questionedmark,zoba,madèle,Kakimasuka,_belo,dragonkid21,felipeherrera7211,Shuun,xXtenthdegree450Xx,benorth,Dingus_Art,RyanHart<3,Yusopp,ZNoteTaku,Mehaniki,deshxun,MetalliGuy2004,marikuga,cindy2465y,Legite,Bah,kuu,8stch,Arial,Tigreladd,Passer,Pascal_Wilson,mango_21,eddiecat,kingberg,kudokana,BadAnimator2024,_Rojas_,carlosg.,EpiclyYYY,poipop_,Arita,operationskuld,TheMr.PrinceKnight,Akajin,Eggnogseller,Sitrep,mrgarci10,Dfrank5547,BigTesty,inSpace,ReverendClown,estier,Linex_anims,zabs,itsagreatdayout,CrispyStar,gestureflow,Jex,ThoyHeeMunn,GabrielRocha,blookx,Chaden,zaris,Gobhit,squaresky89,virgildempsey,kokopia,Obino_Anims,corleonis,Braiynn,dummy,saladboinoli,mjh,apollo,Dylote,d4c,Mokoo,ageingdorito,Oscar_DaSilva,kasshi,Valarie,Nome,RyukoProp,nemfound,lsxxx,Biicriis,ovinha,JDMManga,neptune432,kanon,axxlg,Brian,Troodash,v4msee,minmindu,K.C.pantspants,Enomono,Minie,myammm,Midael,Shreeder4092,veggiemistress,Fat,eggheadmikey,eriiccccc,EpiclyY,cGobongo,meililor,Mr.Seahorse,Fish-Ken,yuki_suzuri,hdd22,Allegedraost,nicaragua,Conanka,crepio,Salademiteuse,rafer,osmansado,HarryVG,skaci,thegroundissky,LeMp,micro,Posty,Tzur,kikishigi,mbuzi,jdeg,mangoamor,しらす,miles1234,mantidk,rzyg,Hydre,VPFL4JZJP5,OSFICC9Y4K,TWJJDM6MTI,Ceayzy,YNMM,DruMzTV,zay,owo,William,Cypriz,cty1228A,yuuex,Gobaa,jem_hiatus,duasmega10,BeastKing,Minhdragon2000,Clemence,Makumo,gobeyond,Ronanig,Lemo,dvon1224,520xinzhao,Silvanno,Shibabibas,mikasa,AdmistNone,MonroeHigh,MekaAoratos,dogmaster,rostrek,Ravenniijima,Eliodesu,krillinkrillin,PurpleGeth,R0S3,MVD-farblo,LKR,Daitayo,TheRealGlitch.EXE,Oscura,N4ssim,saradaba,Wawad_17,angel500x,Gibson,av.tur,Mish,dough,InMotion,Naoufel,grardox,yagoiaba,Spacedad,Yuricide,side_character,yoyoishere,yosihsu,Yoru7,Crest1,Negative,chiearlymorning,silverview,Guancho,MasterCard,TempoHouse,LoogiBaloogi,nav,Ragnite,wormmaster,Gavern,FacuuAF,SALCUL.,xertox,Archaeic,rinze,Grumo,tOothpasTe,ofpveteran73,Pencils,luxlady,Lagressif,Boomer,hiyo,reillu69,aony,Yiyeidi,lalaland4life,BoredomBot,BabyGotBlack890,AkihiraBaka,ender50,astreetninja,BannedUser6313,shankylefou,homen,kye,Hydronic,LHPN,Moonkey,Yuugen_rb,onecoin,FREEDO0,Yessir,tomgeez,Stupidartpunk,Shaymincar,ericolizz,Metal_Oak,Bitcloudyo,airborneandrogynes,katana,FOXAR,Aisahaha4X7,heyitsmojo,FRIIID0O,ThatMangaka,ghmlasco,ken,OmegaLux,kkk,tt_taniel,GVlash,Chao,パンヅ,leonz,Dzian,Clutch,FBarra,5262amag@,Javon,Ozy,Mr.K,Pissed,amin,Charly76,evandro_pedro06,AxB,Eora,SpaRkofFiRe,elmelena,Marc1996,lasersinger,Armando,pokll,pmanrohan,abrasion,gracedotpng,gurdim,infernic,LucasOdoctor,Smil,Bronxden,ShigiDA,hellwinkNG,YaBoiDillPickle,fatbro,KenOzu,darkspike90,漫乃,veka,Hector412,Insight,riooo,Shadow,Thatguy,Theoo,KamKKF,Mysto,ChickenThunderHorse,Zapilaze,Yonan,Mr_Sandman,Rhiannon,NotSally,maujf,zst.u,fypha,kght,dbzfan20039064,Cagliostro,Keins,Adogeandthedoge,Div,naterg,dpm16,NeonDynamite,NCRMN,tuna,ADdictedHD,Tukilit,Doctor-D,Hounddude,Horsebeast,Wsiegmikan,Oracion,protanomania,MrGrim,Gem,Emperor,robin,JazzMazz,jaramo,heyodogo,trough,Machina,Edman,drawphildraw,dgoghvan,SASMf_1122,prokr,TheZeroOne,Nightshade21,Yusk,MatildaAjan,nicodoll,DazVanDamme,molvr,Krisoyo,soy,Tubbsii,Obsidium,Hollingshed,Crash02,osamamii,Yuzuri,akurahiharuka,snappedfrost,tommy73,DT_cool,lucky10aki,potatocakes,Mars,FierceAlchemist,Tashy,francisyfl,skade,Osmel,Miccool,ehlboy,木籽鱼,dragonhunteriv,EternalReverie,JanEstra"
}
`)
