# Hash funkcijos generatorius

## Funkcijos reikalavimai
Sukurti maišos (hash) funkciją, kuri atitiktų šiuos reikalavimus:

Maišos funkcijos įėjimas gali būti bet kokio dydžio simbolių eilutė.
Maišos funkcijos išėjimas visuomet yra to paties, fiksuoto, dydžio rezultatas (pageidautina 256 bit'ų ilgio, t.y., 64 simbolių hex'as).
Maišos funkcija yra deterministinė, t. y., tam pačiam įvedimui išvedimas visuomet yra tas pats.
Maišos funkcijos reikšmė bet kokiai input'o reikšmei yra apskaičiuojamas greitai ir efektyviai.
Iš hash funkcijos rezultato praktiškai neįmanoma atgaminti pradinio įvedimo.
Maišos funkcija yra atspari "kolizijai" (collision resistance), t.y., praktiškai neįmanoma surasti tokių dviejų skirtingų argumentų, kad jiems gautume tą patį hash'ą.
Bent minimaliai pakeitus įvedimą, pvz.vietoj "Lietuva" pateikus "lietuva", maišos funkcijos rezultatas turi skirtis iš esmės, t.y., tenkinamas taip vadinamas lavinos efektas (Avalanche effect).

## Kaip naudotis programa

```bash
go build ./ // Build the program
./hash

Usage of ./hash:
  -comp
        do you want to check for collisions
  -diff
        do you want to find binary/hex difference
  -f string
        provide file for hashing
  -gen
        does this need generating data
  -kon
        do you want algorithm speed comparison
```
## Funkcijos testavimas 

Du failai su vienu bet skirtingu simboliu
```
./hash -f test_files/symbol_b.txt
e2f2309c36fef4186aea98747eb61cb0726280cc46eec4c8fa5ae8a48ea6ec60

./hash -f test_files/symbol_a.txt
baa2b8fc6e0edcd8025ae0947686c430ca9288acfe7e2c08124ab04406f61460
```

Du failai iš daug (>1000) simbolių 
```
./hash -f test_files/random_a.txt    
fef0355486462748e06ae4045feca9109ce4413076f651d07084ade00d849be0

./hash -f test_files/random_b.txt
d4e0c9e043a8995806c81b48bb222c90a03676acb4e02c005556deb8f8785720
```

Du identiški failai kurie skiriasi tik vienu simboliu
```
./hash -f test_files/one_letter_a.txt                               
c6c461f494de2df872089d447eaa1c409a26b168066c0dd819ec623443d87ba0

./hash -f test_files/one_letter_b.txt
0cdaf8d45c0a54e8ec7a18ac78fadaf044ce68acd0d664188af2b27c5206a820
```

Tuščio failo
```
./hash -f test_files/empty.txt
1b64db80535483e06b240b2063d473403b64bb40f3d4e3208b24ebe00354d380
```

Funkcijos greitis lyginant su SHA-256 ir MD5
```
./hash -kon                      
My hash function average is 681 microseconds
SHA256 average is 242 microseconds
MD5 average is 147 microseconds
```

## Pastabos
Tikrinant 100000 eilučių porų, kurios skiriasi tik vienu simboliu, 
kai kuriose iš porų 'identifikuojama' kolizija, nors rankiniu būdu 
paleidus hash funkciją su šiomis eilutėmis, rezultatas nėra tas pats. Šiuomet neatradau jokio 'fix' šitam dalykui, 
bet to priežastimi gali būti per didelis failas (181.4MiB) arba per didelis eilučių kiekis, kas sukelė failų skaitytuve `bufio/Scanner`
UB (Undefined behaviour)

Programos kūrimo metu, kolizijos tikrinimas naudojant `-comp` flag'ą, buvo dukart rasta kolizija, kas yra neišvengiama nekomplikuotame hash algoritme.

## Išvada

Funkcija atitinka visus užduoties reikalavimus, jeigu žinoma, kad kiekvienoje hash funkcijoje galime rasti koliziją.

Funkcija visąlaik gražina 256 bitų outputą

Funkcija yra letesnė nei 128 bitų MD5 ir 256 bitų SHA-256, to priežastimi yra funkcijos primityvumas ir neoptimizuoti skaičiavimai.

Funkcija atspari lavinos (avalanche) efektui. Vieno simbolio skirtumas sudaro vidutiniškai 30% skirtingumą binariniame kode ir 74% skirtingumą šešioliktainėje formoje.

