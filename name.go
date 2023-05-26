package wizlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Name struct {
	First  string
	Middle string
	Last   string
}

type AcceptedNames struct {
	Names []string `json:"names"`
}

// NameGenerator provides methods for generating valid names based on the input and the accepted names list.
type NameGenerator struct {
	acceptedNames AcceptedNames
}

// NewNameGenerator creates a new instance of NameGenerator and retrieves the default accepted names from the provided URL.
func NewNameGenerator(repo NameRepository) (*NameGenerator, error) {
	names, err := repo.GetNames()
	if err != nil {
		return nil, err
	}
	return &NameGenerator{
		acceptedNames: names,
	}, nil
}

// NameRepository defines the contract for accessing name data.
type NameRepository interface {
	GetNames() (AcceptedNames, error)
}

// JSONNameRepository is an implementation of the NameRepository using a JSON file.
type JSONNameRepository struct {
	FilePath string
}

// GetNames retrieves the accepted names from a JSON file.
func (r *JSONNameRepository) GetNames() (AcceptedNames, error) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		return AcceptedNames{}, err
	}
	defer file.Close()

	var names AcceptedNames
	err = json.NewDecoder(file).Decode(&names)
	if err != nil && err != io.EOF {
		return AcceptedNames{}, err
	}

	return names, nil
}

// URLNameRepository is an implementation of the NameRepository using a remote URL.
type URLNameRepository struct {
	URL string
}

// GetNames retrieves the accepted names from a remote URL.
func (r *URLNameRepository) GetNames() (AcceptedNames, error) {
	client := NewAPIClient()
	body, err := client.Get(r.URL)
	if err != nil {
		return AcceptedNames{}, err
	}

	var names AcceptedNames
	err = json.Unmarshal(body, &names)
	if err != nil {
		return AcceptedNames{}, err
	}

	return names, nil
}

// GenerateName generates a valid name based on the input and the accepted names list.
func (g *NameGenerator) GenerateName(input string) (string, error) {
	pattern := fmt.Sprintf(`(?i)^(%s)( (%s))?((%s))?$`, strings.Join(g.acceptedNames.Names, "|"), strings.Join(g.acceptedNames.Names, "|"), strings.Join(g.acceptedNames.Names, "|"))
	nameRegex := regexp.MustCompile(pattern)

	nameParts := strings.Split(input, " ")
	var name Name
	switch len(nameParts) {
	case 1:
		name.First = nameParts[0]
	case 2:
		name.First = nameParts[0]
		name.Last = nameParts[1]
	default:
		name.First = nameParts[0]
		name.Middle = strings.Join(nameParts[1:len(nameParts)-1], " ")
		name.Last = nameParts[len(nameParts)-1]
	}

	if nameRegex.MatchString(input) {
		var sb strings.Builder
		sb.WriteString(name.First)
		if name.Middle != "" {
			sb.WriteString(" ")
			sb.WriteString(name.Middle)
		}
		if name.Last != "" {
			sb.WriteString(name.Last)
		}
		return sb.String(), nil
	}

	return "", errors.New("sorry, the name is not accepted")
}

const defaultNames = `Abigail,Alexandra,Alexandria,Alexis,Alia,Alicia,Allison,Alura,Alyssa,Amanda,Amber,Amy,Andrea,Angela,Ashley,Autumn,Bailey,Brecken,Brianna,Brittany,Brooke,Brynn,Caitlin,Calamity,Caroline,Cassandra,Catherine,Chelsea,Cheryl,Cheyenne,Christina,Cori,Courtney,Danielle,Darby,Deirdre,Delaney,Destiny,Devin,Diana,Donna,Elizabeth,Ellie,Emily,Emma,Emmaline,Erica,Erin,Esmee,Fallon,Fiona,Gabrielle,Genevieve,Ginelle,Grace,Haley,Hannah,Heather,Iridian,Isabella,Jacqueline,Jasmine,Jenna,Jennifer,Jessica,Jordan,Julia,Kaitlyn,Katherine,Katie,Kayla,Keena,Keira,Kelly,Kelsey,Kestrel,Kiley,Kimberly,Kristen,Kymma,Laura,Lauren,Leesha,Lenora,Lindsey,Llewella,Mackenzie,Madeline,Madison,Maria,Mariah,Marissa,Mary,Megan,Melissa,Michelle,Mindy,Miranda,Moria,Molly,Monica,Morgan,Myrna,Natalie,Neela,Nicole,Nora,Olivia,Paige,Rachel,Rebecca,Roslyn,Rowan,Ryan,Rylee,Sabrina,Saffron,Samantha,Sarah,Sarai,Savannah,Scarlet,Sestiva,Shanna,Shannon,Shawna,Shelby,Sierra,Sophia,Stephanie,Suri,Sydney,Tabitha,Taryn,Tasha,Tatiana,Tavia,Terri,Tiffany,Vanessa,Victoria,Aaron,Adam,Adrian,Aedan,Alejandro,Alex,Alexander,Allan,Alric,Andrew,Angel,Angus,Anthony,Antonio,Arlen,Artur,Austin,Belgrim,Benjamin,Blaine,Blake,Blaze,Boris,Bradley,Brady,Brahm,Brand,Brandon,Brian,Caleb,Caley,Cameron,Carlos,Cass,Charles,Chase,Chris,Christo,Christopher,Cody,Cole,Colin,Connor,Corwin,Cowan,Coyle,Dakota,Daniel,Digby,Dolan,Dugan,Duncan,Dustin,Dylan,Edward,Elie,Elijah,Eric,Ethan,Evan,Finnigan,Flint,Fred,Gabriel,Galen,Garret,Gavin,Gilroy,Gorman,Hunter,Ian,Isaac,Isaiah,Jack,Jacob,James,Jared,Jason,Jeffery,Jeremy,Jesse,John,Jonathan,Jose,Joseph,Joshua,Juan,Justin,Kane,Karic,Keelan,Keller,Kenneth,Kevin,Kieran,Kyle,Lail,Liam,Logan,Lucas,Luis,Luke,Malorn,Malvin,Marcus,Mark,Mason,Matthew,Michael,Miguel,Michell,Morgrim,Mycin,Nathan,Nathaniel,Nicholas,Noah,Oran,Padric,Patrick,Paul,Quinn,Reed,Richward,Robert,Rogan,Ronan,Samuel,Scot,Sean,Seth,Sloan,Stephen,Steven,Talon,Tanner,Tarlac,Taylor,Thomas,Timothy,Travis,Trevor,Tristan,Tyler,Valdus,Valerian,Valkoor,William,Wolf,Zachary,Angle,Anvil,Ash,Battle,Bear,Blue,Boom,Crow,Daisy,Dark,Dawn,Day,Death,Dragon,Drake,Dream,Dune,Dusk,Earth,Emerald,Fairy,Fire,Foe,Frog,Frost,Ghost,Giant,Gold,Golden,Green,Griffin,Hawk,Hex,Ice,Iron,Jade,Legend,Life,Light,Lion,Lotus,Mist,Moon,Myth,Night,Ogre,Owl,Pearl,Pixie,Rain,Rainbow,Raven,Red,Rose,Ruby,Sand,Sea,Shadow,Silver,Sky,Skull,Soul,Sparkle,Spell,Spirit,Sprite,Star,Storm,Story,Strong,Summer,Sun,Swift,Tale,Thunder,Titan,Troll,Unicorn,Water,Wild,Willow,Wind,Winter,Wyrm,Bane,Blade,Blood,Bloom,Blossom,Breaker,Breath,Breeze,Bright,Bringer,Caller,Caster,Catcher,Cloud,Coin,Crafter,Dreamer,Dust,Eyes,Finder,Fist,Flame,Flower,Forge,Fountain,Friend,Garden,Gem,Giver,Glade,Glen,Grove,Hammer,Hand,Haven,Head,Heart,Horn,Leaf,Mancer,Mask,Mender,Pants,Petal,Pyre,Rider,River,Runner,Shade,Shard,Shield,Singer,Slinger,Smith,Song,Spear,Staff,Stalker,Steed,Stone,Strider,Sword,Tail,Tamer,Thief,Thistle,Thorn,Vault,Walker,Ward,Weave,Weaver,Whisper,Wielder,Wraith`

func (g *NameGenerator) GetDefaultNames() []string {
	return strings.Split(defaultNames, ",")
}
