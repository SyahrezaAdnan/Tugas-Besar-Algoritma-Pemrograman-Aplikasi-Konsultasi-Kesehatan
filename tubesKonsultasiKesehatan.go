package main

import (
	"fmt"
	"strings"
)

const NMAX = 100

type userKonsultasi struct{
	nama string
	status string
	username string
	email string
	password string
}

type konsultasiPasien struct{
	userYangBertanya userKonsultasi
	userYangMenjawab[NMAX]userKonsultasi
	pertanyaan string
	jawaban[NMAX]string
	dataTag[NMAX]string
	nUserPenjawabDanJawaban int
	nDataTag int
	jumlahKarakter int
}

type tag struct{
	nama string
	jumlahPertanyaan int
}

type totalPertanyaanSesuaiTag[NMAX]tag

type dataPasien struct{
	listPasien [NMAX] userKonsultasi
	nPasien int
}

type dataDokter struct{
	listDokter [NMAX] userKonsultasi
	nDokter int
}

type dataKonsultasi struct{
	listKonsultasi [NMAX] konsultasiPasien
	nKonsultasi int
}

type statusLogin struct{
	apakahLogin bool
	statusUser string
	userYangLogin userKonsultasi
	apakahInginKeluar bool
}


func inisiasiUser(tabDokter *dataDokter, tabPasien *dataPasien){
	// inisiasi tabDokter.listDokter[0]
	tabDokter.listDokter[0].nama = "Syauqi Dhiya Ulhaq"
	tabDokter.listDokter[0].status = "Dokter"
	tabDokter.listDokter[0].username = "syauqi.awigwog"
	tabDokter.listDokter[0].email = "syauqi@gmail.com"
	tabDokter.listDokter[0].password = "12345678"
	tabDokter.nDokter++
	// inisiasi tabDokter.listDokter[1]
	tabDokter.listDokter[1].nama = "Sheriff Clint"
	tabDokter.listDokter[1].status = "Dokter"
	tabDokter.listDokter[1].username = "q"
	tabDokter.listDokter[1].email = "sheriff.clint@gmail.com"
	tabDokter.listDokter[1].password = "w"
	tabDokter.nDokter++
	// inisiasi tabPasien.listPasien[0]
	tabPasien.listPasien[0].nama = "Syahreza Adnan Al Azhar"
	tabPasien.listPasien[0].status = "Pasien"
	tabPasien.listPasien[0].username = "syahreza.adnan"
	tabPasien.listPasien[0].email = "syahreza@gmail.com"
	tabPasien.listPasien[0].password = "AdnanAlAzhar"
	tabPasien.nPasien++
	// inisiasi tabPasien.listPasien[1]
	tabPasien.listPasien[1].nama = "Anton Toni Tino"
	tabPasien.listPasien[1].status = "Pasien"
	tabPasien.listPasien[1].username = "Toni.Tino"
	tabPasien.listPasien[1].email = "ATT@gmail.com"
	tabPasien.listPasien[1].password = "qwertyuiop"
	tabPasien.nPasien++
	// inisiasi tabPasien.listPasien[2]
	tabPasien.listPasien[2].nama = "King Alucard"
	tabPasien.listPasien[2].status = "Pasien"
	tabPasien.listPasien[2].username = "a"
	tabPasien.listPasien[2].email = "king.alucard@gmail.com"
	tabPasien.listPasien[2].password = "s"
	tabPasien.nPasien++
}

func inisiasiJumlahTag(tabTag *totalPertanyaanSesuaiTag){
    tabTag[0].nama = "obat"
    tabTag[0].jumlahPertanyaan = 0
    tabTag[1].nama = "batuk"
    tabTag[1].jumlahPertanyaan = 0
    tabTag[2].nama = "flu"
    tabTag[2].jumlahPertanyaan = 0
    tabTag[3].nama = "radang"
    tabTag[3].jumlahPertanyaan = 0
    tabTag[4].nama = "maag"
    tabTag[4].jumlahPertanyaan = 0
    tabTag[5].nama = "demam"
    tabTag[5].jumlahPertanyaan = 0
    tabTag[6].nama = "diare"
    tabTag[6].jumlahPertanyaan = 0
    tabTag[7].nama = "muntah"
    tabTag[7].jumlahPertanyaan = 0
    tabTag[8].nama = "mual"
    tabTag[8].jumlahPertanyaan = 0
    tabTag[9].nama = "pusing"
    tabTag[9].jumlahPertanyaan = 0
    tabTag[10].nama = "tidak diketahui"
    tabTag[10].jumlahPertanyaan = 0
}

func inisiasiKonsultasi(tabDataKonsultasi *dataKonsultasi, tabTag *totalPertanyaanSesuaiTag, tabDokter *dataDokter, tabPasien *dataPasien) {
	// inisiasi tabDataKonsultasi.listKonsultasi[0]
	tabDataKonsultasi.listKonsultasi[0].userYangBertanya = tabPasien.listPasien[0]
	tabDataKonsultasi.listKonsultasi[0].pertanyaan = "apakah kalau sakit batuk harus minum obat?"

	tabDataKonsultasi.listKonsultasi[0].userYangMenjawab[tabDataKonsultasi.listKonsultasi[0].nUserPenjawabDanJawaban] = tabDokter.listDokter[0]
	tabDataKonsultasi.listKonsultasi[0].jawaban[tabDataKonsultasi.listKonsultasi[0].nUserPenjawabDanJawaban] = "tidak harus minum obat, bisa dengan minum banyak air putih."
	tabDataKonsultasi.listKonsultasi[0].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[0].userYangMenjawab[tabDataKonsultasi.listKonsultasi[0].nUserPenjawabDanJawaban] = tabPasien.listPasien[1]
	tabDataKonsultasi.listKonsultasi[0].jawaban[tabDataKonsultasi.listKonsultasi[0].nUserPenjawabDanJawaban] = "harus minum obat yang sangat banyak"
	tabDataKonsultasi.listKonsultasi[0].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[0].dataTag[tabDataKonsultasi.listKonsultasi[0].nDataTag] = "obat"
	tabDataKonsultasi.listKonsultasi[0].nDataTag++
	tabDataKonsultasi.listKonsultasi[0].dataTag[tabDataKonsultasi.listKonsultasi[0].nDataTag] = "batuk"
	tabDataKonsultasi.listKonsultasi[0].nDataTag++

	tabTag[0].jumlahPertanyaan++
	tabTag[1].jumlahPertanyaan++

	tabDataKonsultasi.listKonsultasi[0].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[0].pertanyaan)

	tabDataKonsultasi.nKonsultasi++

	// inisiasi tabDataKonsultasi.listKonsultasi[1]
	tabDataKonsultasi.listKonsultasi[1].userYangBertanya = tabPasien.listPasien[1]
	tabDataKonsultasi.listKonsultasi[1].pertanyaan = "Saya sakit apa ya dok? Sudah tiga hari pusing, mual, dan muntah."

	tabDataKonsultasi.listKonsultasi[1].dataTag[tabDataKonsultasi.listKonsultasi[1].nDataTag] = "pusing"
	tabDataKonsultasi.listKonsultasi[1].nDataTag++
	tabDataKonsultasi.listKonsultasi[1].dataTag[tabDataKonsultasi.listKonsultasi[1].nDataTag] = "mual"
	tabDataKonsultasi.listKonsultasi[1].nDataTag++
	tabDataKonsultasi.listKonsultasi[1].dataTag[tabDataKonsultasi.listKonsultasi[1].nDataTag] = "muntah"
	tabDataKonsultasi.listKonsultasi[1].nDataTag++

	tabDataKonsultasi.listKonsultasi[1].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[1].pertanyaan)

	tabTag[9].jumlahPertanyaan++
	tabTag[8].jumlahPertanyaan++
	tabTag[7].jumlahPertanyaan++

	tabDataKonsultasi.nKonsultasi++

	// inisiasi tabDataKonsultasi.listKonsultasi[2]
	tabDataKonsultasi.listKonsultasi[2].userYangBertanya = tabPasien.listPasien[0]
	tabDataKonsultasi.listKonsultasi[2].pertanyaan = "apakah radang menyebabkan batuk?"

	tabDataKonsultasi.listKonsultasi[2].userYangMenjawab[tabDataKonsultasi.listKonsultasi[2].nUserPenjawabDanJawaban] = tabDokter.listDokter[0]
	tabDataKonsultasi.listKonsultasi[2].jawaban[tabDataKonsultasi.listKonsultasi[2].nUserPenjawabDanJawaban] = "Jika terjadi peradangan pada faring, muncul gejala berupa gatal di tenggorokan. Rasa gatal tersebut yang memicu terjadinya batuk."
	tabDataKonsultasi.listKonsultasi[2].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[2].dataTag[tabDataKonsultasi.listKonsultasi[2].nDataTag] = "radang"
	tabDataKonsultasi.listKonsultasi[2].nDataTag++
	tabDataKonsultasi.listKonsultasi[2].dataTag[tabDataKonsultasi.listKonsultasi[2].nDataTag] = "batuk"
	tabDataKonsultasi.listKonsultasi[2].nDataTag++

	tabTag[3].jumlahPertanyaan++
	tabTag[1].jumlahPertanyaan++

	tabDataKonsultasi.listKonsultasi[2].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[2].pertanyaan)

	tabDataKonsultasi.nKonsultasi++

	// inisiasi tabDataKonsultasi.listKonsultasi[3]
	tabDataKonsultasi.listKonsultasi[3].userYangBertanya = tabPasien.listPasien[1]
	tabDataKonsultasi.listKonsultasi[3].pertanyaan = "apakah diare mematikan?"

	tabDataKonsultasi.listKonsultasi[3].userYangMenjawab[tabDataKonsultasi.listKonsultasi[3].nUserPenjawabDanJawaban] = tabDokter.listDokter[0]
	tabDataKonsultasi.listKonsultasi[3].jawaban[tabDataKonsultasi.listKonsultasi[3].nUserPenjawabDanJawaban] = "."
	tabDataKonsultasi.listKonsultasi[3].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[3].dataTag[tabDataKonsultasi.listKonsultasi[3].nDataTag] = "diare"
	tabDataKonsultasi.listKonsultasi[3].nDataTag++

	tabTag[6].jumlahPertanyaan++

	tabDataKonsultasi.listKonsultasi[3].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[3].pertanyaan)

	tabDataKonsultasi.nKonsultasi++

	// inisiasi tabDataKonsultasi.listKonsultasi[4]
	tabDataKonsultasi.listKonsultasi[4].userYangBertanya = tabPasien.listPasien[2]
	tabDataKonsultasi.listKonsultasi[4].pertanyaan = "Apakah flu dapat menyebabkan demam juga?"

	tabDataKonsultasi.listKonsultasi[4].userYangMenjawab[tabDataKonsultasi.listKonsultasi[4].nUserPenjawabDanJawaban] = tabDokter.listDokter[1]
	tabDataKonsultasi.listKonsultasi[4].jawaban[tabDataKonsultasi.listKonsultasi[4].nUserPenjawabDanJawaban] = "Ya, flu berpontensi untuk meningkatkan suhu tubuh"
	tabDataKonsultasi.listKonsultasi[4].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[4].dataTag[tabDataKonsultasi.listKonsultasi[4].nDataTag] = "flu"
	tabDataKonsultasi.listKonsultasi[4].nDataTag++
	tabDataKonsultasi.listKonsultasi[4].dataTag[tabDataKonsultasi.listKonsultasi[4].nDataTag] = "demam"
	tabDataKonsultasi.listKonsultasi[4].nDataTag++

	tabTag[2].jumlahPertanyaan++
	tabTag[5].jumlahPertanyaan++

	tabDataKonsultasi.listKonsultasi[4].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[4].pertanyaan)

	tabDataKonsultasi.nKonsultasi++

	// inisiasi tabDataKonsultasi.listKonsultasi[5]
	tabDataKonsultasi.listKonsultasi[5].userYangBertanya = tabPasien.listPasien[0]
	tabDataKonsultasi.listKonsultasi[5].pertanyaan = "Apakah hipermetropi dapat ditimbulkan dari menonton tv terlalu dekat?"

	tabDataKonsultasi.listKonsultasi[5].userYangMenjawab[tabDataKonsultasi.listKonsultasi[5].nUserPenjawabDanJawaban] = tabDokter.listDokter[0]
	tabDataKonsultasi.listKonsultasi[5].jawaban[tabDataKonsultasi.listKonsultasi[5].nUserPenjawabDanJawaban] = "Mitos, hanya saja ketika menonton terlalu dekat mata akan tegang dan perlu diistirahatkan"
	tabDataKonsultasi.listKonsultasi[5].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[5].dataTag[tabDataKonsultasi.listKonsultasi[5].nDataTag] = "tidak diketahui"
	tabDataKonsultasi.listKonsultasi[5].nDataTag++

	tabTag[10].jumlahPertanyaan++

	tabDataKonsultasi.listKonsultasi[5].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[5].pertanyaan)

	tabDataKonsultasi.nKonsultasi++

	// inisiasi tabDataKonsultasi.listKonsultasi[6]
	tabDataKonsultasi.listKonsultasi[6].userYangBertanya = tabPasien.listPasien[1]
	tabDataKonsultasi.listKonsultasi[6].pertanyaan = "Saya suka pusing dok, apakah saya vertigo?"

	tabDataKonsultasi.listKonsultasi[6].userYangMenjawab[tabDataKonsultasi.listKonsultasi[6].nUserPenjawabDanJawaban] = tabDokter.listDokter[1]
	tabDataKonsultasi.listKonsultasi[6].jawaban[tabDataKonsultasi.listKonsultasi[6].nUserPenjawabDanJawaban] = "Tidak, bisa jadi anda kecapean"
	tabDataKonsultasi.listKonsultasi[6].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[6].dataTag[tabDataKonsultasi.listKonsultasi[6].nDataTag] = "pusing"
	tabDataKonsultasi.listKonsultasi[6].nDataTag++

	tabTag[9].jumlahPertanyaan++

	tabDataKonsultasi.listKonsultasi[6].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[6].pertanyaan)

	tabDataKonsultasi.nKonsultasi++

	// inisiasi tabDataKonsultasi.listKonsultasi[7]
	tabDataKonsultasi.listKonsultasi[7].userYangBertanya = tabPasien.listPasien[2]
	tabDataKonsultasi.listKonsultasi[7].pertanyaan = "Maag saya sering kambuh dok, apakah bisa disembuhkan?"

	tabDataKonsultasi.listKonsultasi[7].userYangMenjawab[tabDataKonsultasi.listKonsultasi[7].nUserPenjawabDanJawaban] = tabDokter.listDokter[0]
	tabDataKonsultasi.listKonsultasi[7].jawaban[tabDataKonsultasi.listKonsultasi[7].nUserPenjawabDanJawaban] = "Ya, minum promag dan jaga selalu jam makan anda"
	tabDataKonsultasi.listKonsultasi[7].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[7].dataTag[tabDataKonsultasi.listKonsultasi[7].nDataTag] = "maag"
	tabDataKonsultasi.listKonsultasi[7].nDataTag++

	tabTag[4].jumlahPertanyaan++

	tabDataKonsultasi.listKonsultasi[7].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[7].pertanyaan)

	tabDataKonsultasi.nKonsultasi++

	// inisiasi tabDataKonsultasi.listKonsultasi[8]
	tabDataKonsultasi.listKonsultasi[8].userYangBertanya = tabPasien.listPasien[0]
	tabDataKonsultasi.listKonsultasi[8].pertanyaan = "Apakah gejala covid itu diare dok?"

	tabDataKonsultasi.listKonsultasi[8].userYangMenjawab[tabDataKonsultasi.listKonsultasi[8].nUserPenjawabDanJawaban] = tabDokter.listDokter[1]
	tabDataKonsultasi.listKonsultasi[8].jawaban[tabDataKonsultasi.listKonsultasi[8].nUserPenjawabDanJawaban] = "Terinfeksi covid bisa menyebabkan diare"
	tabDataKonsultasi.listKonsultasi[8].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[8].dataTag[tabDataKonsultasi.listKonsultasi[8].nDataTag] = "diare"
	tabDataKonsultasi.listKonsultasi[8].nDataTag++

	tabTag[6].jumlahPertanyaan++

	tabDataKonsultasi.listKonsultasi[8].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[8].pertanyaan)

	tabDataKonsultasi.nKonsultasi++

	// inisiasi tabDataKonsultasi.listKonsultasi[9]
	tabDataKonsultasi.listKonsultasi[9].userYangBertanya = tabPasien.listPasien[1]
	tabDataKonsultasi.listKonsultasi[9].pertanyaan = "Apakah obat amoxicillin bisa menyembuhkan radang dok?"

	tabDataKonsultasi.listKonsultasi[9].userYangMenjawab[tabDataKonsultasi.listKonsultasi[9].nUserPenjawabDanJawaban] = tabDokter.listDokter[1]
	tabDataKonsultasi.listKonsultasi[9].jawaban[tabDataKonsultasi.listKonsultasi[9].nUserPenjawabDanJawaban] = "Ya, karena obat itu dapat menyembuhkan infeksi bakteri"
	tabDataKonsultasi.listKonsultasi[9].nUserPenjawabDanJawaban++

	tabDataKonsultasi.listKonsultasi[9].dataTag[tabDataKonsultasi.listKonsultasi[9].nDataTag] = "obat"
	tabDataKonsultasi.listKonsultasi[9].nDataTag++
	tabDataKonsultasi.listKonsultasi[9].dataTag[tabDataKonsultasi.listKonsultasi[9].nDataTag] = "radang"
	tabDataKonsultasi.listKonsultasi[9].nDataTag++

	tabTag[0].jumlahPertanyaan++
	tabTag[3].jumlahPertanyaan++

	tabDataKonsultasi.listKonsultasi[9].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[9].pertanyaan)

	tabDataKonsultasi.nKonsultasi++
}

func inisiasiStatusLoginAwal(login *statusLogin){
	login.apakahLogin = false
	login.statusUser = ""
	login.userYangLogin.nama = ""
	login.userYangLogin.status = ""
	login.userYangLogin.username = ""
	login.userYangLogin.email = ""
	login.userYangLogin.password = ""
	login.apakahInginKeluar = false
}

func adaUsernameYangSama(tabDokter dataDokter, tabPasien dataPasien, inputUsername string) int {
	//melakukan pengecekan inputUsername apakah ada username yang sama pada tabDokter atau tabPasien
	for i := 0; i < tabDokter.nDokter; i++{
		if tabDokter.listDokter[i].username == inputUsername{
			return 1000000+i
		}
	}
	for i := 0; i < tabPasien.nPasien; i++{
		if tabPasien.listPasien[i].username == inputUsername{
			return 2000000+i
		}
	}
	return 0
}

func inputWithSpace(totalString *string){
	
	var stringPanjang, inputKata string
	stringPanjang = ""
	fmt.Scan(&inputKata)
	for inputKata != "STOP"{
		stringPanjang += inputKata + " "
		fmt.Scan(&inputKata)
	}
	*totalString = stringPanjang
}

func register(tabDokter *dataDokter, tabPasien *dataPasien){
	var inputStatus string
	var userNameSama int
	fmt.Println()
	fmt.Println("==================================")
	fmt.Println("============Register==============")
	fmt.Println("==================================")
	fmt.Println()
	//melakukan input Status apakah sesuai dengan status yang tersedia
	fmt.Println("Apa status anda? (Dokter/Pasien)")
	fmt.Scan(&inputStatus)
	for inputStatus != "Dokter" && inputStatus != "Pasien"{
		fmt.Println("Status anda salah, status yang tersedia adalah Dokter dan Pasien. Huruf pertama harus menggunakan huruf kapital")
		fmt.Scan(&inputStatus)
	}

	//penyimpanan data user yang baru berdasarkan statusnya
	if inputStatus == "Dokter"{
		//menyimpan data user yang baru pada tabDokter
		tabDokter.listDokter[tabDokter.nDokter].status = "Dokter"
		fmt.Println("Masukkan nama lengkap anda (akhiri nama dengan kata STOP): ")
		inputWithSpace(&tabDokter.listDokter[tabDokter.nDokter].nama)
		fmt.Println("Masukkan username anda: ")
		fmt.Scan(&tabDokter.listDokter[tabDokter.nDokter].username)

		//looping jika username sudah ada sebelumnya
		userNameSama = adaUsernameYangSama(*tabDokter, *tabPasien, tabDokter.listDokter[tabDokter.nDokter].username)
		for userNameSama != 0{
			fmt.Println("Username sudah ada, masukkan username yang lain")
			fmt.Println("Masukkan username yang lain:")
			fmt.Scan(&tabDokter.listDokter[tabDokter.nDokter].username)
			userNameSama = adaUsernameYangSama(*tabDokter, *tabPasien, tabDokter.listDokter[tabDokter.nDokter].username)
		}

		fmt.Println("Masukkan email anda:")
		fmt.Scan(&tabDokter.listDokter[tabDokter.nDokter].email)
		fmt.Println("Masukkan password anda:")
		fmt.Scan(&tabDokter.listDokter[tabDokter.nDokter].password)
		tabDokter.nDokter++

	}else if inputStatus == "Pasien"{
		//menyimpan data user yang baru pada tabPasien
		tabPasien.listPasien[tabPasien.nPasien].status = "Pasien"
		fmt.Println("Masukkan nama lengkap anda (akhiri nama dengan kata STOP): ")
		inputWithSpace(&tabPasien.listPasien[tabPasien.nPasien].nama)
		fmt.Println("Masukkan username anda:")
		fmt.Scan(&tabPasien.listPasien[tabPasien.nPasien].username)

		//melakukan looping jika username sudah ada sebelumnya
		userNameSama = adaUsernameYangSama(*tabDokter, *tabPasien, tabPasien.listPasien[tabPasien.nPasien].username)
		for userNameSama != 0{
			fmt.Println("Username sudah ada, masukkan username yang lain")
			fmt.Println("Masukkan username yang lain:")
			fmt.Scan(&tabPasien.listPasien[tabPasien.nPasien].username)
			userNameSama = adaUsernameYangSama(*tabDokter, *tabPasien, tabPasien.listPasien[tabPasien.nPasien].username)
		}

		fmt.Println("Masukkan email anda:")
		fmt.Scan(&tabPasien.listPasien[tabPasien.nPasien].email)
		fmt.Println("Masukkan password anda:")
		fmt.Scan(&tabPasien.listPasien[tabPasien.nPasien].password)
		tabPasien.nPasien++
	}
}

func login(tabDokter dataDokter, tabPasien dataPasien, dataLogin *statusLogin){
	var inputUsername, inputPassword string
	var idxUsername int
	var percobaanLogin int = 0

	fmt.Println()
	fmt.Println("==================================")
	fmt.Println("=============Log In===============")
	fmt.Println("==================================")
	fmt.Println()

	fmt.Println("Masukkan username anda:")
	fmt.Scan(&inputUsername)
	idxUsername = adaUsernameYangSama(tabDokter, tabPasien, inputUsername)

	//looping jika username tidak ditemukan
	for idxUsername == 0 && percobaanLogin < 3{
		percobaanLogin++
		fmt.Println("Username tidak ditemukan")
		fmt.Println("Masukkan username anda:")
		fmt.Scan(&inputUsername)
		idxUsername = adaUsernameYangSama(tabDokter, tabPasien, inputUsername)
		if percobaanLogin == 3{
			fmt.Println("Anda telah mencoba 3 kali, silahkan registrasi terlebih dahulu")
			register(&tabDokter, &tabPasien)
			fmt.Println("Silahkan login kembali")
			fmt.Print("Masukkan username anda:")
			fmt.Scan(&inputUsername)
			idxUsername = adaUsernameYangSama(tabDokter, tabPasien, inputUsername)
		}
	}

	//penyimpanan username yang login, dan penyimpanan data login, dan melakukan input password
	if idxUsername >= 2000000{
		dataLogin.statusUser = "Pasien"
		dataLogin.apakahLogin = true
		dataLogin.userYangLogin = tabPasien.listPasien[idxUsername-2000000]

		//melakukan proses input password
		fmt.Println("Masukkan password anda:")
		fmt.Scan(&inputPassword)
		for inputPassword != dataLogin.userYangLogin.password{
			fmt.Println("Password salah")
			fmt.Println("Masukkan password anda:")
			fmt.Scan(&inputPassword)
		}
	}else{
		dataLogin.statusUser = "Dokter"
		dataLogin.apakahLogin = true
		dataLogin.userYangLogin = tabDokter.listDokter[idxUsername-1000000]

		//melakukan proses input password
		fmt.Println("Masukkan password anda:")
		fmt.Scan(&inputPassword)
		for inputPassword != dataLogin.userYangLogin.password{
			fmt.Println("Password salah")
			fmt.Println("Masukkan password anda:")
			fmt.Scan(&inputPassword)
		}
	}
	
	fmt.Println("Login berhasil")
	fmt.Print("Selamat datang ", dataLogin.userYangLogin.nama, ", pada forum konsultasi kesehatan SAYHAT")
	fmt.Println()
	fmt.Println("Status anda adalah", dataLogin.userYangLogin.status)
}

func printForumKonsultasi(tabDataKonsultasi dataKonsultasi){
	fmt.Println("==================================")
	fmt.Println("=============SAYHAT===============")
	fmt.Println("====Forum Konsultasi Kesehatan====")
	fmt.Println("==================================")
	fmt.Println()
	fmt.Println()
	for i := 0; i < tabDataKonsultasi.nKonsultasi; i++{
		fmt.Print(i+1,". ", tabDataKonsultasi.listKonsultasi[i].userYangBertanya.nama, " (", tabDataKonsultasi.listKonsultasi[i].userYangBertanya.status, ")")
		fmt.Print("\n")
		fmt.Print("     >> ", tabDataKonsultasi.listKonsultasi[i].pertanyaan)
		fmt.Print("\n")
		fmt.Print("     Tags: ")
		for k := 0; k < tabDataKonsultasi.listKonsultasi[i].nDataTag; k++{
			fmt.Print("#", tabDataKonsultasi.listKonsultasi[i].dataTag[k], " ")
		}
		fmt.Print("\n","\n")
		fmt.Print("     vvv Jawaban vvv")
		fmt.Print("\n", "\n")
		for j := 0; j < tabDataKonsultasi.listKonsultasi[i].nUserPenjawabDanJawaban; j++{
			fmt.Print("     >>> ", j+1, ". ", tabDataKonsultasi.listKonsultasi[i].userYangMenjawab[j].nama, " (", tabDataKonsultasi.listKonsultasi[i].userYangMenjawab[j].status, ")")
			fmt.Print("\n")
			fmt.Print("     >>>>> ", tabDataKonsultasi.listKonsultasi[i].jawaban[j])
			fmt.Print("\n")
			fmt.Print("     --------------------", "\n")
		}

		if tabDataKonsultasi.listKonsultasi[i].nUserPenjawabDanJawaban == 0{
			fmt.Print("     Belum ada jawaban", "\n")
		}

		fmt.Print("\n", "\n")
	}
}

func inputPertanyaan(tabDokter *dataDokter, tabPasien *dataPasien, dataLogin *statusLogin, tabDataKonsultasi *dataKonsultasi, tabTag *totalPertanyaanSesuaiTag){
	var input1atau2 int
	if dataLogin.apakahLogin{
		fmt.Println("Masukkan pertanyaan anda (akhiri pertanyaan dengan kata STOP): ")
		inputWithSpace(&tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].pertanyaan)
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].userYangBertanya = dataLogin.userYangLogin
		pemberianTag(tabDataKonsultasi, tabTag)
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].jumlahKarakter = menghitungPanjangPertanyaan(tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].pertanyaan)
		tabDataKonsultasi.nKonsultasi++
		fmt.Println()
		fmt.Println("========Pertanyaan anda berhasil ditambahkan===========")
		fmt.Println("=======================================================")
		fmt.Println()
		fmt.Println()
		printForumKonsultasi(*tabDataKonsultasi)
	}else{
		fmt.Println("Harap log in terlebih dahulu")
		fmt.Println("Ketik 1 untuk registrasi")
		fmt.Println("Ketik 2 untuk log in")
		fmt.Println("============================")
		fmt.Scan(&input1atau2)
		if input1atau2 == 1{
			register(tabDokter, tabPasien)
		}else if input1atau2 == 2{
			login(*tabDokter, *tabPasien, dataLogin)
		}
	}
}

func hasilSearchingForumKonsultasiSesuaiDenganTag(tabDataKonsultasi dataKonsultasi, dataTag totalPertanyaanSesuaiTag, tag string){
	var tempI int = 0

	for m := 0 ; m < 11; m++{
		if dataTag[m].nama == tag{
			if dataTag[m].jumlahPertanyaan != 0{
				fmt.Print("Terdapat ", dataTag[m].jumlahPertanyaan, " pertanyaan pada tag #", tag, "\n")
				fmt.Print("Berikut list pertanyaan yang sesuai dengan tag #", tag,"\n", "\n")
			}
		}
	}

	for i := 0; i < tabDataKonsultasi.nKonsultasi; i++{

		for j := 0; j < tabDataKonsultasi.listKonsultasi[i].nDataTag; j++{
			if tabDataKonsultasi.listKonsultasi[i].dataTag[j] == tag{
				fmt.Print(tempI+1,". ", tabDataKonsultasi.listKonsultasi[i].userYangBertanya.nama, " (", tabDataKonsultasi.listKonsultasi[i].userYangBertanya.status, ")")
				fmt.Print("\n")
				fmt.Print("     >> ", tabDataKonsultasi.listKonsultasi[i].pertanyaan)
				fmt.Print("\n")
				fmt.Print("     Tags: ")
				for k := 0; k < tabDataKonsultasi.listKonsultasi[i].nDataTag; k++{
					fmt.Print("#", tabDataKonsultasi.listKonsultasi[i].dataTag[k], " ")
				}
				fmt.Print("\n","\n")
				fmt.Print("     vvv Jawaban vvv")
				fmt.Print("\n", "\n")
				for l := 0; l < tabDataKonsultasi.listKonsultasi[i].nUserPenjawabDanJawaban; l++{
					fmt.Print("     >>> ", l+1, ". ", tabDataKonsultasi.listKonsultasi[i].userYangMenjawab[l].nama, " (", tabDataKonsultasi.listKonsultasi[i].userYangMenjawab[l].status, ")")
					fmt.Print("\n")
					fmt.Print("     >>>>> ", tabDataKonsultasi.listKonsultasi[i].jawaban[l])
					fmt.Print("\n")
					fmt.Print("     --------------------", "\n")
				}

				if tabDataKonsultasi.listKonsultasi[i].nUserPenjawabDanJawaban == 0{
					fmt.Print("     Belum ada jawaban", "\n")
				}

				fmt.Print("\n", "\n")
				tempI++
			}
		}
	}

	for m := 0 ; m < 11; m++{
		if dataTag[m].nama == tag{
			if dataTag[m].jumlahPertanyaan == 0{
				fmt.Print("Tidak ada pertanyaan yang sesuai dengan tag #", tag, "\n", "\n")
			}
		}
	}
	fmt.Scanln()
}

func binarySearchTag(tabTag totalPertanyaanSesuaiTag, intPencarian int) int{
	var found int = -1
	var med int
	var kr int = 0
	var kn int = len(tabTag)-1
	for kr <= kn && found == -1{
		med = (kr+kn)/2
		if intPencarian > tabTag[med].jumlahPertanyaan{
			kn = med - 1
		}else if intPencarian < tabTag[med].jumlahPertanyaan{
			kr = med + 1
		}else{
			found = med
		}
	}
	return found
}

func hasilSearchingForumKonsultasiSesuaiDenganJumlahPertanyaanPadaTag(tabDataKonsultasi dataKonsultasi, tabTag totalPertanyaanSesuaiTag, intPencarian int){
	var tempTabTag totalPertanyaanSesuaiTag
	var penemuanTag[NMAX]tag
	var nPenemuanTag int = 0
	var idxBinary int
	var nTempTabTag int = 10
	tempTabTag = tabTag
	for i := 0; i <= 10; i++{
		idxBinary = binarySearchTag(tempTabTag, intPencarian)
		if idxBinary != -1{
			penemuanTag[i] = tempTabTag[idxBinary]
			nPenemuanTag = i+1
			for j := idxBinary; j < nTempTabTag; j++ {
				tempTabTag[j] = tempTabTag[j+1]
			}
			nTempTabTag--
		}
	}
	for r := 0; r < nPenemuanTag; r++{
		nomor := 0
		fmt.Print("Hasil pencarian dengan jumlah pertanyaan pada tag #", penemuanTag[r].nama, ":", "\n")
		for i := 0; i < tabDataKonsultasi.nKonsultasi; i++{
			for j := 0; j < tabDataKonsultasi.listKonsultasi[i].nDataTag; j++{
				if tabDataKonsultasi.listKonsultasi[i].dataTag[j] == penemuanTag[r].nama{
		
					fmt.Print(nomor+1,". ", tabDataKonsultasi.listKonsultasi[i].userYangBertanya.nama, " (", tabDataKonsultasi.listKonsultasi[i].userYangBertanya.status, ")")
					nomor++
					fmt.Print("\n")
					fmt.Print("     >> ", tabDataKonsultasi.listKonsultasi[i].pertanyaan)
					fmt.Print("\n")
					fmt.Print("     Tags: ")
					for k := 0; k < tabDataKonsultasi.listKonsultasi[i].nDataTag; k++{
						fmt.Print("#", tabDataKonsultasi.listKonsultasi[i].dataTag[k], " ")
					}
					fmt.Print("\n","\n")
					fmt.Print("     vvv Jawaban vvv")
					fmt.Print("\n", "\n")
					for l := 0; l < tabDataKonsultasi.listKonsultasi[i].nUserPenjawabDanJawaban; l++{
						fmt.Print("     >>> ", l+1, ". ", tabDataKonsultasi.listKonsultasi[i].userYangMenjawab[l].nama, " (", tabDataKonsultasi.listKonsultasi[i].userYangMenjawab[l].status, ")")
						fmt.Print("\n")
						fmt.Print("     >>>>> ", tabDataKonsultasi.listKonsultasi[i].jawaban[l])
						fmt.Print("\n")
						fmt.Print("     --------------------", "\n")
					}
	
					if tabDataKonsultasi.listKonsultasi[i].nUserPenjawabDanJawaban == 0{
						fmt.Print("     Belum ada jawaban", "\n")
					}
	
					fmt.Print("\n", "\n")
				}
			}
		}
		fmt.Println("=======================================================")
	}
}

func inputJawaban(tabDokter *dataDokter, tabPasien *dataPasien, tabDataKonsultasi *dataKonsultasi, dataLogin *statusLogin){
	var inputNomorPertanyaan int
	var inputJawaban string
	printForumKonsultasi(*tabDataKonsultasi)
	fmt.Println("Masukkan nomor pertanyaan yang ingin anda jawab:")
	fmt.Scan(&inputNomorPertanyaan)
	for inputNomorPertanyaan < 1 || inputNomorPertanyaan > tabDataKonsultasi.nKonsultasi{
		fmt.Println("Nomor pertanyaan yang anda masukkan salah")
		fmt.Println("Masukkan nomor pertanyaan yang ingin anda jawab:")
		fmt.Scan(&inputNomorPertanyaan)
	}
	fmt.Println("Masukkan jawaban anda (akhiri jawaban dengan kata STOP):")
	inputWithSpace(&inputJawaban)
	tabDataKonsultasi.listKonsultasi[inputNomorPertanyaan-1].userYangMenjawab[tabDataKonsultasi.listKonsultasi[inputNomorPertanyaan-1].nUserPenjawabDanJawaban] = dataLogin.userYangLogin
	tabDataKonsultasi.listKonsultasi[inputNomorPertanyaan-1].jawaban[tabDataKonsultasi.listKonsultasi[inputNomorPertanyaan-1].nUserPenjawabDanJawaban] = inputJawaban
	tabDataKonsultasi.listKonsultasi[inputNomorPertanyaan-1].nUserPenjawabDanJawaban++
	fmt.Println("Jawaban anda berhasil ditambahkan")
}

func menuUtama(tabDokter *dataDokter, tabPasien *dataPasien, dataLogin *statusLogin, tabDataKonsultasi *dataKonsultasi, tabTag *totalPertanyaanSesuaiTag) {
	var input int
	for input != 4 && !dataLogin.apakahInginKeluar{
		if !dataLogin.apakahInginKeluar{
			fmt.Println("==================================")
			fmt.Println("========     SAYHAT     ==========")
			fmt.Println("========   Menu Utama   ==========")
			fmt.Println("====Forum Konsultasi Kesehatan====")
			fmt.Println("==================================")
			fmt.Println()
			fmt.Println()
			fmt.Println("Pilih Menu:")
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("3. Forum Konsultasi")
			fmt.Println("4. Keluar dari SAYHAT")
			fmt.Println("Masukkan nomor pilihan menu:")
			fmt.Scan(&input)
			for input != 1 && input != 2 && input != 3 && input != 4 {
				fmt.Println("Nomor menu yang anda masukkan salah")
				fmt.Println("Pilih nomer menu kembali:")
				fmt.Scan(&input)
			}
		}
		if input == 1 {
			register(tabDokter, tabPasien)
			fmt.Println("Anda harus log in terlebih dahulu untuk dapat mengakses menu Forum Konsultasi SAYHAT")
			login(*tabDokter, *tabPasien, dataLogin)
			menuForumKonsultasi(tabDataKonsultasi, dataLogin, tabDokter, tabPasien, tabTag)
		} else if input == 2 {
			login(*tabDokter, *tabPasien, dataLogin)
			menuForumKonsultasi(tabDataKonsultasi, dataLogin, tabDokter, tabPasien, tabTag)
		} else if input == 3 {
			menuForumKonsultasi(tabDataKonsultasi, dataLogin, tabDokter, tabPasien, tabTag)
		} else if input == 4 {
			dataLogin.apakahInginKeluar = true
			fmt.Println()
			fmt.Println("Terima kasih telah menggunakan SAYHAT")
			fmt.Println("Sampai jumpa")
			fmt.Println("Jika ada pertanyaan, bertanya di SAYHAT saja ya :D")
		}
	}
}

func menuForumKonsultasi(tabDataKonsultasi *dataKonsultasi, dataLogin *statusLogin, tabDokter *dataDokter, tabPasien *dataPasien, tabTag *totalPertanyaanSesuaiTag) {
	var input int
	if dataLogin.apakahLogin {
		for input != 5 && !dataLogin.apakahInginKeluar{
			if !dataLogin.apakahInginKeluar{
				fmt.Println("========================")
				fmt.Println("    Forum Konsultasi")
				fmt.Println("========================")
				fmt.Println("1. Tampilkan Forum Konsultasi")
				fmt.Println("2. Bertanya")
				fmt.Println("3. Mencari Pertanyaan")
				fmt.Println("4. Menjawab")
				fmt.Println("5. Kembali ke Menu Utama")
				fmt.Println("Pilih nomer menu:")
				fmt.Scan(&input)
				for input != 1 && input != 2 && input != 3 && input != 4 && input != 5{
					fmt.Println("Nomor menu yang anda masukkan salah")
					fmt.Println("Pilih nomer menu kembali:")
					fmt.Scan(&input)
				}
			}
			if input == 1{
				printForumKonsultasi(*tabDataKonsultasi)
			}else if input == 2 {
				inputPertanyaan(tabDokter, tabPasien, dataLogin, tabDataKonsultasi, tabTag)
			} else if input == 3 {
				if dataLogin.statusUser == "Dokter" {
					menuTagDokter(tabDataKonsultasi, *tabTag, dataLogin, tabDokter, tabPasien)
				}else if dataLogin.statusUser == "Pasien" {
					menuTagPasien(tabDataKonsultasi, *tabTag, dataLogin)
				}
			} else if input == 4 {
				inputJawaban(tabDokter, tabPasien, tabDataKonsultasi, dataLogin)
			} else if input == 5 { 
				menuUtama(tabDokter, tabPasien, dataLogin, tabDataKonsultasi, tabTag)
			}
		}
	} else {
		for input != 5{
			if !dataLogin.apakahInginKeluar{
				fmt.Println("========================")
				fmt.Println("    Forum Konsultasi")
				fmt.Println("========================")
				fmt.Println("1. Tampilkan Forum Konsultasi")
				fmt.Println("2. Bertanya")
				fmt.Println("3. Mencari Pertanyaan")
				fmt.Println("4. Menjawab")
				fmt.Println("5. Kembali ke Menu Utama")
				fmt.Println("Pilih Menu:")
				fmt.Scan(&input)
				for input != 1 && input != 2 && input != 3 && input != 4 && input != 5{
					fmt.Println("Nomor menu yang anda masukkan salah")
					fmt.Println("Pilih nomer menu kembali:")
					fmt.Scan(&input)
				}
			}
			if input == 1{
				printForumKonsultasi(*tabDataKonsultasi)
			}else if input == 2 || input == 4 {
				fmt.Println("Anda harus Log In terlebih dahulu!")
				fmt.Println("1. Register")
				fmt.Println("2. Log In")
				fmt.Println("3. Kembali ke Menu")
				fmt.Scan(&input)
				for input != 1 && input != 2 && input != 3 {
					fmt.Println("Nomor menu yang anda masukkan salah")
					fmt.Println("Pilih nomer menu kembali:")
					fmt.Scan(&input)
				}
				if input == 1{
					register(tabDokter, tabPasien)
					fmt.Println("Anda harus log in terlebih dahulu untuk dapat mengakses menu Forum Konsultasi SAYHAT")
					login(*tabDokter, *tabPasien, dataLogin)
				}else if input == 2 {
					login(*tabDokter, *tabPasien, dataLogin)
				}else if input == 3 {
					menuForumKonsultasi(tabDataKonsultasi, dataLogin, tabDokter, tabPasien, tabTag)
				}
			} else if input == 3 {
				menuTagPasien(tabDataKonsultasi, *tabTag, dataLogin)
			} else if input == 5 {
				menuUtama(tabDokter, tabPasien, dataLogin, tabDataKonsultasi, tabTag)
			}
		}
	}
}

func menghitungPanjangPertanyaan(pertanyaan string) int {
	return len(pertanyaan)
}

func mencariSubstring(pertanyaan, tag string) bool {
	return strings.Contains(pertanyaan, tag)
}

func ToLower(s string) string {
	var lower string
	for _, c := range s {
	   if c >= 'A' && c <= 'Z' {
		  lower += string(c + 32)
	   } else {
		  lower += string(c)
	   }
	}
	return lower
 }

func pemberianTag(tabDataKonsultasi *dataKonsultasi, tabTag *totalPertanyaanSesuaiTag) {
	var tempPertanyaan string = ToLower(tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].pertanyaan)
	if mencariSubstring(tempPertanyaan, "obat") {
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "obat"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "obat" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	} 
	if mencariSubstring(tempPertanyaan, "batuk") {
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "batuk"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "batuk" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	} 
	if mencariSubstring(tempPertanyaan, "flu") {
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "flu"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "flu" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	} 
	if mencariSubstring(tempPertanyaan, "radang"){
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "radang"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "radang" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	} 
	if mencariSubstring(tempPertanyaan, "maag"){
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "maag"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "maag" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	}
	if mencariSubstring(tempPertanyaan, "demam"){
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "demam"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "demam" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	} 
	if mencariSubstring(tempPertanyaan, "diare"){
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "diare"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "diare" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	}
	if mencariSubstring(tempPertanyaan, "muntah"){
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "muntah"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "muntah" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	}
	if mencariSubstring(tempPertanyaan, "mual"){
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "mual"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "mual" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	}
	if mencariSubstring(tempPertanyaan, "pusing"){
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "pusing"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "pusing" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	}
	if tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag == 0 {
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].dataTag[tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag] = "tidak diketahui"
		tabDataKonsultasi.listKonsultasi[tabDataKonsultasi.nKonsultasi].nDataTag++
		for i := 0; i <= 10; i++ {
			if tabTag[i].nama == "tidak diketahui" {
				tabTag[i].jumlahPertanyaan++
			}
		}
	}
}

func sortingForumKonsultasi(tabTag *totalPertanyaanSesuaiTag) {
	var temp tag
	var i, j, idxMAX int
	i = 1
	for i < 11 {
		idxMAX = i - 1
		j = i
		for j < 11 {
			if tabTag[idxMAX].jumlahPertanyaan < tabTag[j].jumlahPertanyaan {
				idxMAX = j
			}
			j = j + 1
		}
		temp = tabTag[idxMAX]
		tabTag[idxMAX] = tabTag[i-1]
		tabTag[i-1] = temp
		i = i + 1
	}
}

func sortingPertanyaan(tabDataKonsultasi *dataKonsultasi) {
	var temp konsultasiPasien
	var i, pass int
	pass = 1
	for pass < tabDataKonsultasi.nKonsultasi {
		i = pass
		temp = tabDataKonsultasi.listKonsultasi[i]
		for i > 0 && temp.jumlahKarakter > tabDataKonsultasi.listKonsultasi[i-1].jumlahKarakter {
			tabDataKonsultasi.listKonsultasi[i] = tabDataKonsultasi.listKonsultasi[i-1]
			i--
		}
		tabDataKonsultasi.listKonsultasi[i] = temp
		pass++
	}
}

func menuTagPasien(tabDataKonsultasi *dataKonsultasi, tabTag totalPertanyaanSesuaiTag, dataLogin *statusLogin) {
	var input int
	if !dataLogin.apakahInginKeluar {
		fmt.Println("==========================")
		fmt.Println("       Daftar Tags")
		fmt.Println("==========================")
		for i := 0; i <= 10; i++ {
			fmt.Print(i+1)
			if i < 9 {
				fmt.Print(".  ")
			} else {
				fmt.Print(". ")
			}
			fmt.Print("#", tabTag[i].nama,"\n")
		}
		fmt.Println("Pilih nomor tag:")
		fmt.Scan(&input)
		for input != 1 && input != 2 && input != 3 && input != 4 && input != 5 && input != 6 && input != 7 && input != 8 && input != 9 && input != 10 && input != 11 {
			fmt.Println("Nomor tag yang anda masukkan salah")
			fmt.Println("Pilih nomor tag kembali:")
			fmt.Scan(&input)
		}
	}
	if input == 1 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[0].nama)
	} else if input == 2 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[1].nama)
	} else if input == 3 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[2].nama)
	} else if input == 4 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[3].nama)
	} else if input == 5 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[4].nama)
	} else if input == 6 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[5].nama)
	} else if input == 7 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[6].nama)
	} else if input == 8 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[7].nama)
	} else if input == 9 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[8].nama)
	} else if input == 10 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[9].nama)
	} else if input == 11 {
		hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[10].nama)
	}
}

func menuTagDokter(tabDataKonsultasi *dataKonsultasi, tabTag totalPertanyaanSesuaiTag, dataLogin *statusLogin, tabDokter *dataDokter, tabPasien *dataPasien){
	var input, inputJumlahPertanyaan int
	var tempTabDataKonsultasi dataKonsultasi
	tempTabDataKonsultasi = *tabDataKonsultasi
	for input != 4 && !dataLogin.apakahInginKeluar{
		if !dataLogin.apakahInginKeluar{
			fmt.Println("==========================")
			fmt.Println("       Menu Pencarian")
			fmt.Println("==========================")
			fmt.Println("1. Mencari Bedasarkan Nama Tag")
			fmt.Println("2. Mencari Bedasarkan Jumlah Pertanyaan pada Tag")
			fmt.Println("3. Tampilkan Pertanyaan dari Pertanyaan Terpanjang")
			fmt.Println("4. Kembali")
			fmt.Println("Pilih nomor metode pencarian:")
			fmt.Scan(&input)
			for input != 1 && input != 2 && input != 3 && input != 4 {
				fmt.Println("Nomor metode pencarian yang anda masukkan salah")
				fmt.Println("Pilih nomor metode pencarian kembali:")
				fmt.Scan(&input)
			}
		}
		if input == 1 {
			sortingForumKonsultasi(&tabTag)
			fmt.Println("==========================")
			fmt.Println("       Daftar Tags")
			fmt.Println("==========================")
			for i := 0; i <= 10; i++ {
				fmt.Print(i+1)
				if i < 9 {
					fmt.Print(".  ")
				} else {
					fmt.Print(". ")
				}
				fmt.Print("#",tabTag[i].nama)
				for j := 0; j <= 20 - len(tabTag[i].nama); j++{
					fmt.Print(" ")
				}
				fmt.Print("(", tabTag[i].jumlahPertanyaan, " pertanyaan)", "\n")
			}
			fmt.Println()
			fmt.Println("Pilih nomor tag:")
			fmt.Scan(&input)
			for input != 1 && input != 2 && input != 3 && input != 4 && input != 5 && input != 6 && input != 7 && input != 8 && input != 9 && input != 10 && input != 11 {
				fmt.Println("Nomor tag yang anda masukkan salah")
				fmt.Println("Pilih nomor tag kembali:")
				fmt.Scan(&input)
			}
			if input == 1 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[0].nama)
			} else if input == 2 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[1].nama)
			} else if input == 3 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[2].nama)
			} else if input == 4 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[3].nama)
			} else if input == 5 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[4].nama)
			} else if input == 6 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[5].nama)
			} else if input == 7 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[6].nama)
			} else if input == 8 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[7].nama)
			} else if input == 9 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[8].nama)
			} else if input == 10 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[9].nama)
			} else if input == 11 {
				hasilSearchingForumKonsultasiSesuaiDenganTag(*tabDataKonsultasi, tabTag, tabTag[10].nama)
			}
		} else if input == 2 {
			sortingForumKonsultasi(&tabTag)
			fmt.Println("==========================")
			fmt.Println("       Daftar Tags")
			fmt.Println("==========================")
			for i := 0; i <= 10; i++ {
				fmt.Print(i+1)
				if i < 9 {
					fmt.Print(".  ")
				} else {
					fmt.Print(". ")
				}
				fmt.Print("#",tabTag[i].nama)
				for j := 0; j <= 20 - len(tabTag[i].nama); j++{
					fmt.Print(" ")
				}
				fmt.Print("(", tabTag[i].jumlahPertanyaan, " pertanyaan)", "\n")
			}
			fmt.Println("==========================")
			fmt.Println("Masukkan jumlah pertanyaan yang ingin dicari sesuai dengan tag:")
			fmt.Scan(&inputJumlahPertanyaan)
			fmt.Println()
			fmt.Println("Hasil pencarian:")
			fmt.Println("===========================================")
			hasilSearchingForumKonsultasiSesuaiDenganJumlahPertanyaanPadaTag(*tabDataKonsultasi, tabTag, inputJumlahPertanyaan)
		} else if input == 3 {
			sortingPertanyaan(&tempTabDataKonsultasi)
			printForumKonsultasi(*&tempTabDataKonsultasi)
		} else if input == 4{
			menuForumKonsultasi(tabDataKonsultasi, dataLogin, tabDokter, tabPasien, &tabTag)
		}
	}
}

func main(){
	var tabDataKonsultasi dataKonsultasi
	var tabTag totalPertanyaanSesuaiTag
	var tabDokter dataDokter
	var tabPasien dataPasien
	var dataLogin statusLogin
	inisiasiUser(&tabDokter, &tabPasien)
	inisiasiJumlahTag(&tabTag)
	inisiasiKonsultasi(&tabDataKonsultasi, &tabTag, &tabDokter, &tabPasien)
	inisiasiStatusLoginAwal(&dataLogin)
	menuUtama(&tabDokter, &tabPasien, &dataLogin, &tabDataKonsultasi, &tabTag)
	
}