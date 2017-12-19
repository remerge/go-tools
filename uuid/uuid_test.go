package uuid

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var uuids = []struct {
	id            string
	possiblyValid bool
	valid         bool
}{
	{"asdads", false, false},
	{"0X000000A0000Z0000Y00Y0X000000000000", false, false},
	{"00000000-0000-0000-0000-000000000000", true, false},
	{"0X000000-0000-0000-00Y0-000000000000", true, false},
	{"FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF", true, false},
	{"ffffffff-ffff-ffff-ffff-ffffffffffff", true, false},
	{"eD83F96D-D14E-4A81-8C7A-021DA055DEF5", true, true},
	{"a4128ee3-5fd5-4359-bae4-6c03354f9a02", true, true},
	{"28C1E1D8-7E76-4C57-A25A-4B8C81A3EB2A", true, true},
	{"eD83F96D-D14E-4A81-8C7A-021DA055DEF5F", false, false},
}

func TestUUID(t *testing.T) {
	for _, uuid := range uuids {
		possiblyValid := isValidFast([]byte(uuid.id))
		if uuid.possiblyValid != possiblyValid {
			t.Errorf("%s should be recognized as possibltyValid=%t but is possibltyValid=%t", uuid.id, uuid.possiblyValid, possiblyValid)
		}
		valid := IsValid([]byte(uuid.id))
		if uuid.valid != valid {
			fmt.Printf("%s should be recognized as valid=%t but is valid=%t\n", uuid.id, uuid.valid, valid)
			t.Errorf("%s should be recognized as valid=%t but is valid=%t", uuid.id, uuid.valid, valid)
		}
	}
}

func TestUUIDOsMatching(t *testing.T) {
	assert.True(t, MatchesOS("", "ios", []byte("bla")), "an empty os should always match")
	assert.False(t, MatchesOS("android", "ios", []byte("bla")), "different OSs shouldn't match")
	assert.False(t, MatchesOS("android", "android", []byte("ED83F96D-D14E-4A81-8C7A-021DA055DEF5")), "correct OS but wrong user id shouldn't match")
	assert.True(t, MatchesOS("ios", "ios", []byte("ED83F96D-D14E-4A81-8C7A-021DA055DEF5")), "correct OS but correct user id should match")
}

var uuidRegex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

var benchUuids = [][]byte{
	[]byte("04506430-f5c1-4d19-a54e-d5cee6a9355c"),
	[]byte("0a4c9008-a367-4e15-885f-87a2872b813b"),
	[]byte("10742e2f-9c41-43c0-b467-f5accc82853a"),
	[]byte("109b4f20-75dd-4515-85ad-68c53ae83c80"),
	[]byte("15e34bda-ce0f-42b6-9f98-38eb020f3f6e"),
	[]byte("25be2cc9-3d13-402f-9022-8999fb3956d7"),
	[]byte("30a8b3ad-d0af-4511-8b89-823e05b8dc68"),
	[]byte("33144d0c-8962-4978-9063-39b496250273"),
	[]byte("438a5cd6-2e48-49bf-bf0a-adea96e9b6ae"),
	[]byte("48051443-5a57-48b8-83a0-437e5c0b3848"),
	[]byte("5853619a-950e-4d72-b5d0-48607a5abfe1"),
	[]byte("75b8378f-5b3c-4631-87e8-157472c93654"),
	[]byte("7b6d12a8-5248-451a-b493-906720120853"),
	[]byte("9dc7ebd5-cd30-470a-bedc-8a2ee44e8765"),
	[]byte("Being7igheefeixaiteu4aeba"),
	[]byte("Bip8eetoo7Eu4ieng4uyethae"),
	[]byte("ShieCh4mait2quu1eez6eisoo"),
	[]byte("UDoh5oraeQu0lohf3eecho3oo"),
	[]byte("WeeWee7oopoogeif6eLohShae"),
	[]byte("a40a1ab1-8a24-4fb7-ae3a-f8fef16d3020"),
	[]byte("a78b7769-1052-498f-a9bb-c06de3aafd2d"),
	[]byte("aK9oocah9quaniech3shoo1bo0dip4fuojiequ9A"),
	[]byte("aivahmiYofo8quiwaahahquooy4ohPh8joi5aiZa"),
	[]byte("apheiDa8fei6oongeutoophahCo3Ohve0uPh"),
	[]byte("b0d4c216-2d16-4ff7-b309-5018532e3b3d"),
	[]byte("b3cd49e7-52cb-41e1-8c84-82f476fc1928"),
	[]byte("baoVeiZaeshe9iezethaeKee5if1Piequ7iedoV1"),
	[]byte("bdf47899-90d1-4055-8970-7a312661369c"),
	[]byte("bootah6ohf4goh0aecioSijo8aiyae4ahNgahj6E"),
	[]byte("c42d0cfe-12a2-4718-b68e-6b805c308463"),
	[]byte("cab4pi9vietie4uaj8quiTelu"),
	[]byte("cd2d9cc6-c20c-4286-b8d7-cb6c64495095"),
	[]byte("coolo6ohtheeXa6neu8AeJeh5ju6aePh6xai"),
	[]byte("d7b9fec1-4c30-49c0-bb26-99f6a28a53e7"),
	[]byte("db600736-0677-4778-9861-f549b8906ba6"),
	[]byte("epheiso3kieL5iel5aujuixei6tohzahc8oonico"),
	[]byte("equoo8eeloohae1eigoo5eYer"),
	[]byte("f7df06e1-82ac-4e8b-bb0d-cf92241f019d"),
	[]byte("fe1a0709-997e-48b8-80a0-272fb6d268d5"),
	[]byte("feih0loom0hooSh1cae8Baesae7Ohgoo2aj5"),
	[]byte("gas8Goyai1soh5aivoh7eeSh1bi4cheiXoh6"),
	[]byte("geiyio2Aish9yee1Heam4eija1Iejaicuituotoo"),
	[]byte("iang3Eegairo9SaeMooNi8cah9ootaiw7Beaghet"),
	[]byte("machee4iShailahkei4shooGhohlaeweehoh"),
	[]byte("mei1Wa5feiGh0eeNgiey2aong"),
	[]byte("ohph1ot0Euce8IeghaiPhaidiiToohaeb0ph"),
	[]byte("ooghahN1vaezoQueesh4airooxiNgoo3obohhoh4"),
	[]byte("ooyae9earo8quoh3EeseiTh5v"),
	[]byte("phu6Faepoepahc2oht2choj0za4ful0jeepohz2I"),
	[]byte("quohwoh6yooghoy6ech4tis6eiWee7utohFo"),
	[]byte("waih2vo9iofo0aiLeehaehievu5aB5ooquo9eing"),
	[]byte("wei1zo5thoh9voodiasu5Queej0gaeX4aiko"),
	[]byte("xaoChiuj0aetabair0weeth7e"),
	[]byte("ye6Oiba2coh7eFidiilothuach4aep1taphe"),
	[]byte("zeik3Eeth9veiPaewahzaeteiC9samauween"),
}

func BenchmarkUuidRegex(b *testing.B) {
	var hits int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uuid := benchUuids[hits%len(benchUuids)]
			hits += 1
			uuidRegex.Match(uuid)
		}
	})
}

func BenchmarkUuidRagel(b *testing.B) {
	var hits int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uuid := benchUuids[hits%len(benchUuids)]
			hits += 1
			matchUuidRegex(uuid)
		}
	})
}
