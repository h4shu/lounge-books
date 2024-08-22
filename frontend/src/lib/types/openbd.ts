export type Root = Root2[];

export interface Root2 {
	onix: Onix;
	hanmoto: Hanmoto;
	summary: Summary;
}

export interface Onix {
	RecordReference: string;
	NotificationType: string;
	ProductIdentifier: ProductIdentifier;
	DescriptiveDetail: DescriptiveDetail;
	CollateralDetail: CollateralDetail;
	PublishingDetail: PublishingDetail;
	ProductSupply: ProductSupply;
}

export interface ProductIdentifier {
	ProductIDType: string;
	IDValue: string;
}

export interface DescriptiveDetail {
	ProductComposition: string;
	ProductForm: string;
	ProductFormDetail: string;
	TitleDetail: TitleDetail;
	Contributor: Contributor[];
	Language: Language[];
	Extent: Extent[];
	Subject: Subject[];
	Audience: Audience[];
}

export interface TitleDetail {
	TitleType: string;
	TitleElement: TitleElement;
}

export interface TitleElement {
	TitleElementLevel: string;
	TitleText: TitleText;
	Subtitle: Subtitle;
}

export interface TitleText {
	collationkey: string;
	content: string;
}

export interface Subtitle {
	collationkey: string;
	content: string;
}

export interface Contributor {
	SequenceNumber: string;
	ContributorRole: string[];
	PersonName: PersonName;
	BiographicalNote?: string;
}

export interface PersonName {
	collationkey: string;
	content: string;
}

export interface Language {
	LanguageRole: string;
	LanguageCode: string;
	CountryCode: string;
}

export interface Extent {
	ExtentType: string;
	ExtentValue: string;
	ExtentUnit: string;
}

export interface Subject {
	MainSubject?: string;
	SubjectSchemeIdentifier: string;
	SubjectCode: string;
}

export interface Audience {
	AudienceCodeType: string;
	AudienceCodeValue: string;
}

export interface CollateralDetail {
	TextContent: TextContent[];
	SupportingResource: SupportingResource[];
}

export interface TextContent {
	TextType: string;
	ContentAudience: string;
	Text: string;
}

export interface SupportingResource {
	ResourceContentType: string;
	ContentAudience: string;
	ResourceMode: string;
	ResourceVersion: ResourceVersion[];
}

export interface ResourceVersion {
	ResourceForm: string;
	ResourceVersionFeature: ResourceVersionFeature[];
	ResourceLink: string;
}

export interface ResourceVersionFeature {
	ResourceVersionFeatureType: string;
	FeatureValue: string;
}

export interface PublishingDetail {
	Imprint: Imprint;
	Publisher: Publisher;
	PublishingDate: PublishingDate[];
}

export interface Imprint {
	ImprintIdentifier: ImprintIdentifier[];
	ImprintName: string;
}

export interface ImprintIdentifier {
	ImprintIDType: string;
	IDValue: string;
}

export interface Publisher {
	PublishingRole: string;
	PublisherIdentifier: PublisherIdentifier[];
	PublisherName: string;
}

export interface PublisherIdentifier {
	PublisherIDType: string;
	IDValue: string;
}

export interface PublishingDate {
	PublishingDateRole: string;
	Date: string;
}

export interface ProductSupply {
	MarketPublishingDetail: MarketPublishingDetail;
	SupplyDetail: SupplyDetail;
}

export interface MarketPublishingDetail {
	MarketPublishingStatus: string;
	MarketPublishingStatusNote: string;
}

export interface SupplyDetail {
	ProductAvailability: string;
	Price: Price[];
}

export interface Price {
	PriceType: string;
	PriceAmount: string;
	CurrencyCode: string;
}

export interface Hanmoto {
	genshomei: string;
	han: string;
	datezeppan: string;
	toji: string;
	zaiko: number;
	maegakinado: string;
	kaisetsu105w: string;
	tsuiki: string;
	genrecodetrc: number;
	ndccode: string;
	kankoukeitai: string;
	sonotatokkijikou: string;
	jushoujouhou: string;
	furokusonota: string;
	dokushakakikomi: string;
	zasshicode: string;
	jyuhan: Jyuhan[];
	hatsubai: string;
	hatsubaiyomi: string;
	hastameshiyomi: boolean;
	storelink: string;
	author: Author[];
	datemodified: string;
	datecreated: string;
	kanrenshoisbn: string;
	reviews: Review[];
	hanmotoinfo: Hanmotoinfo;
	hankeidokuji: string;
	dateshuppan: string;
}

export interface Jyuhan {
	date: string;
	ctime: string;
	suri: number;
	comment?: string;
}

export interface Author {
	listseq: number;
	dokujikubun: string;
}

export interface Review {
	post_user: string;
	reviewer: string;
	source_id: number;
	kubun_id: number;
	source: string;
	choyukan: string;
	han: string;
	link: string;
	appearance: string;
	gou: string;
}

export interface Hanmotoinfo {
	name: string;
	yomi: string;
	url: string;
	twitter: string;
	facebook: string;
	chokutori: string;
	toritsugitorikyo: string;
	toritsugisonota: string;
	eigyoudaihyousha: string;
	phoneshoten: string;
	facsimileshoten: string;
	emailshoten: string;
	ordersite: string;
	ordersitesonota: string;
	ordersitejisha: string;
	henpin: string;
	jiyuukinyuu: string;
}

export interface Summary {
	isbn: string;
	title: string;
	volume: string;
	series: string;
	publisher: string;
	pubdate: string;
	cover: string;
	author: string;
}
