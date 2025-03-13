package sakugaapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/stretchr/testify/assert"
)

func TestFavoritesListUsers(t *testing.T) {
	t.Run("base case, happy path", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.RequestURI() == "/favorite/list_users.json?id=3" {
				w.WriteHeader(http.StatusOK)
				w.Write(favoriteUsersData)
			}
		}))
		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)

		opts := sakugamodels.FavoriteListUsersOptions{
			ID: 3,
		}
		response, err := newAPI.Favorites.ListUsers(&opts)
		assert.Nil(t, err)
		expected := sakugamodels.FavoriteListUsersResponse{FavoritedUsers: []string{"Kawaki", "Zzdapao", "Nave", "Jupiterjavelin", "Moondoggie", "Dominic_Falisi", "Anitrix", "Petrychrores", "SandyWick", "mg13", "kaise2347", "m2", "Miyagi", "Ravennnn", "Spicygolf", "Fy3", "DanielSonic", "Stouze", "Mokuga", "Clon0s", "Nic2007", "martingrasoso", "Werdert", "swimhigh",
			"sakinreyy", "lvl2frog", "sakugayballs", "Brubthevideogameduck", "drake366", "mzamecki", "Miyuki", "goatt", "Welipe", "mybigheartforyou", "29dan5", "Jen", "akina432", "poipop_tiktok", "Martin_Caubet", "questionedmark", "zoba", "madèle", "Kakimasuka",
			"_belo", "dragonkid21", "felipeherrera7211", "Shuun", "xXtenthdegree450Xx", "benorth", "Dingus_Art", "RyanHart<3", "Yusopp", "ZNoteTaku", "Mehaniki", "deshxun", "MetalliGuy2004", "marikuga", "cindy2465y", "Legite", "Bah", "kuu", "8stch", "Arial", "Tigreladd", "Passer", "Pascal_Wilson", "mango_21", "eddiecat", "kingberg", "kudokana", "BadAnimator2024",
			"_Rojas_", "carlosg.", "EpiclyYYY", "poipop_", "Arita", "operationskuld", "TheMr.PrinceKnight", "Akajin", "Eggnogseller", "Sitrep", "mrgarci10", "Dfrank5547", "BigTesty", "inSpace", "ReverendClown", "estier", "Linex_anims", "zabs", "itsagreatdayout",
			"CrispyStar", "gestureflow", "Jex", "ThoyHeeMunn", "GabrielRocha", "blookx", "Chaden", "zaris", "Gobhit", "squaresky89", "virgildempsey", "kokopia", "Obino_Anims", "corleonis", "Braiynn", "dummy", "saladboinoli", "mjh", "apollo", "Dylote", "d4c", "Mokoo", "ageingdorito", "Oscar_DaSilva", "kasshi", "Valarie", "Nome", "RyukoProp", "nemfound",
			"lsxxx", "Biicriis", "ovinha", "JDMManga", "neptune432", "kanon", "axxlg", "Brian", "Troodash", "v4msee", "minmindu", "K.C.pantspants", "Enomono", "Minie", "myammm", "Midael", "Shreeder4092", "veggiemistress", "Fat", "eggheadmikey", "eriiccccc", "EpiclyY", "cGobongo", "meililor", "Mr.Seahorse", "Fish-Ken", "yuki_suzuri", "hdd22", "Allegedraost",
			"nicaragua", "Conanka", "crepio", "Salademiteuse", "rafer", "osmansado", "HarryVG", "skaci", "thegroundissky", "LeMp", "micro", "Posty", "Tzur", "kikishigi", "mbuzi", "jdeg", "mangoamor", "しらす", "miles1234", "mantidk", "rzyg", "Hydre", "VPFL4JZJP5",
			"OSFICC9Y4K", "TWJJDM6MTI", "Ceayzy", "YNMM", "DruMzTV", "zay", "owo", "William", "Cypriz", "cty1228A", "yuuex", "Gobaa", "jem_hiatus", "duasmega10", "BeastKing", "Minhdragon2000", "Clemence", "Makumo", "gobeyond", "Ronanig", "Lemo", "dvon1224", "520xinzhao", "Silvanno", "Shibabibas", "mikasa", "AdmistNone", "MonroeHigh", "MekaAoratos",
			"dogmaster", "rostrek", "Ravenniijima", "Eliodesu", "krillinkrillin", "PurpleGeth", "R0S3", "MVD-farblo", "LKR", "Daitayo", "TheRealGlitch.EXE", "Oscura", "N4ssim", "saradaba", "Wawad_17", "angel500x", "Gibson", "av.tur", "Mish", "dough", "InMotion", "Naoufel",
			"grardox", "yagoiaba", "Spacedad", "Yuricide", "side_character", "yoyoishere", "yosihsu", "Yoru7", "Crest1", "Negative", "chiearlymorning", "silverview", "Guancho", "MasterCard", "TempoHouse", "LoogiBaloogi", "nav", "Ragnite", "wormmaster", "Gavern", "FacuuAF", "SALCUL.", "xertox", "Archaeic", "rinze", "Grumo", "tOothpasTe", "ofpveteran73",
			"Pencils", "luxlady", "Lagressif", "Boomer", "hiyo", "reillu69", "aony", "Yiyeidi", "lalaland4life", "BoredomBot", "BabyGotBlack890", "AkihiraBaka", "ender50", "astreetninja", "BannedUser6313", "shankylefou",
			"homen", "kye", "Hydronic", "LHPN", "Moonkey", "Yuugen_rb", "onecoin", "FREEDO0", "Yessir", "tomgeez", "Stupidartpunk", "Shaymincar", "ericolizz", "Metal_Oak", "Bitcloudyo", "airborneandrogynes", "katana", "FOXAR", "Aisahaha4X7", "heyitsmojo", "FRIIID0O", "ThatMangaka", "ghmlasco", "ken", "OmegaLux", "kkk", "tt_taniel", "GVlash", "Chao", "パンヅ",
			"leonz", "Dzian", "Clutch", "FBarra", "5262amag@", "Javon", "Ozy", "Mr.K", "Pissed", "amin", "Charly76", "evandro_pedro06", "AxB", "Eora", "SpaRkofFiRe", "elmelena", "Marc1996", "lasersinger", "Armando", "pokll", "pmanrohan", "abrasion", "gracedotpng",
			"gurdim", "infernic", "LucasOdoctor", "Smil", "Bronxden", "ShigiDA", "hellwinkNG", "YaBoiDillPickle", "fatbro", "KenOzu", "darkspike90", "漫乃", "veka", "Hector412", "Insight", "riooo", "Shadow", "Thatguy", "Theoo", "KamKKF", "Mysto", "ChickenThunderHorse", "Zapilaze", "Yonan", "Mr_Sandman", "Rhiannon", "NotSally", "maujf", "zst.u", "fypha",
			"kght", "dbzfan20039064", "Cagliostro", "Keins", "Adogeandthedoge", "Div", "naterg", "dpm16", "NeonDynamite", "NCRMN", "tuna", "ADdictedHD", "Tukilit", "Doctor-D", "Hounddude", "Horsebeast", "Wsiegmikan", "Oracion", "protanomania", "MrGrim", "Gem", "Emperor",
			"robin", "JazzMazz", "jaramo", "heyodogo", "trough", "Machina", "Edman", "drawphildraw", "dgoghvan", "SASMf_1122", "prokr", "TheZeroOne", "Nightshade21", "Yusk", "MatildaAjan", "nicodoll", "DazVanDamme", "molvr", "Krisoyo", "soy", "Tubbsii", "Obsidium", "Hollingshed", "Crash02", "osamamii", "Yuzuri", "akurahiharuka", "snappedfrost", "tommy73",
			"DT_cool", "lucky10aki", "potatocakes", "Mars", "FierceAlchemist", "Tashy", "francisyfl", "skade", "Osmel", "Miccool", "ehlboy", "木籽鱼", "dragonhunteriv", "EternalReverie", "JanEstra"}}
		assert.Equal(t, expected, response)
	})
}
