import type { ProductsType } from 'myTypes/chat.types';

//CONSTANTS
export const emptyProduct: ProductsType = [
	{
		image_url: '',
		product_name: '',
		description_product: '',
		ean: '',
		flag: undefined,
		rating: Infinity,
		offers: [{ store_name: '', logo_url: '', price: '', address: '', currency: 'PLN' }]
	}
];

export const flags = [
	'Najlepszy wybór',
	'Ekonomiczny wybór',
	'Najtańszy dobrze oceniany',
	'Najlepiej oceniany'
];

// COMMON
export const commonPartStartMessage =
	'C: Jestem  generatorem statusów. ' +
	'O: Chcę stworzyć krótki tekst w formie statusu. ' +
	'S: W stylu profesjonalnym ale nie korporacyjnym. ' +
	'A: Ludzie od 16 roku życia. ' +
	'T: W tonie uprzejmym. ' +
	'R: Odpowiedz w języku, w którym sformułowano zapytanie, zachowując poprawność gramatyczną, że ';

const commonAnswwerStart =
	'C: Jestem profesjonalnym doradcą zakupowym działającym TYLKO w oparciu o dostarczoną listę ofert. ' +
	'O: Chcę stworzyć krótką odpowiedź dla użytkownika. ' +
	'S: W stylu profesjonalnym ale nie korporacyjnym. ' +
	'T: W tonie uprzejmym. ' +
	'R: Odpowiedz w języku, w którym sformułowano zapytanie, zachowując poprawność gramatyczną.';

// STATUS
export const startText =
	commonPartStartMessage + 'weryfikujesz posiadane informacje na podany temat w bazie wiedzy.';

export const helpersTextBeforeCheckingCategoryCheck = (hasHistory: boolean, historyCats: boolean) =>
	`C: Jestem precyzyjnym klasyfikatorem intencji.
		###
    O: Klasyfikacja zapytania użytkownika do jednej z kategorii systemowych. 
    ###
    S: Styl techniczny, jednoznaczny. 
    ###
    A: System logiki aplikacji.
    ###
    R: Odpowiedz TYLKO nazwą intencji (np. CHAT):\n
     - zakaz proponowania kategorii produktów lub wyszukiwania w internecie i sklepach;\n
    - skup się wyłącznie na klasyfikacji do poniższych intencji. Ignoruj zapytania o produkty spoza listy; \n
    - STAN HISTORII: ${hasHistory && historyCats ? 'ISTNIEJE' : 'BRAK'};\n
		- MOŻLIWE INTENCJE:\n` +
	'--- BEZ WZGLĘDU NA STAN HISTORII: ---\n' +
	'* "END_SESSION" -> Użytkownik kończy rozmowę (np. "Dzięki, to wszystko", "Super, kupuję", "Nic mi nie pasuje, kończymy", "Znalazłem produkt");\n' +
	'* "OTHER" -> Luźne pogawędki (np. "Jak leci?", "Co słychać?"), spam lub tematy niezwiązane z zakupami;\n' +
	'--- JEŚLI STAN HISTORII TO "BRAK": ---\n' +
	'* "GREETING" -> Samo przywitanie (np. "Hej", "Dzień dobry") lub przywitanie połączone od razu z potrzebą zakupową (np. "Cześć, szukam butów").\n' +
	'* "CHAT" -> Tylko pierwsze pytanie o konkretny produkt.Może być z przywitaniem np. "Cześć, znajdź mi chłodziarki do win"\n' +
	'--- JEŚLI STAN HISTORII TO "ISTNIEJE": ---\n' +
	'* "SEARCH_INFO" -> Pytanie o cechy, detale lub porównanie produktów z historii;\n' +
	'* "SEARCH_MORE" -> Prośba o więcej propozycji, inne modele lub tańsze alternatywy w tej samej kategorii;\n' +
	'* "SEARCH_NEW" -> Radykalna zmiana kategorii produktów na inną. \n' +
	'- odpowiedz tylko nazwą intencji np. CHAT.';

export const helpersTextBeforeCheckingCategory =
	commonPartStartMessage +
	'- wiadomość nie dotyczy szukania produktów;' +
	'- jesteś wyszkolony TYLKO do wyszukiwania produktów;' +
	'- do innych celów polecasz użycie innych narzędzi.';

export const searchingCategory =
	commonPartStartMessage + 'rozpoczynasz rozpoznawanie kategorii dla tego zapytania.';

export const noCategoryMessage =
	commonPartStartMessage +
	'zapytanie nie pasuje do obecnie wspieranych kategorii produktów' +
	' i zachęć do spróbowania znalezienia innego produktu lub spróbowanie ponownie za jakiś czas.';

export const foundCategories = (categoriesLimited: { slug: string }[]) =>
	commonPartStartMessage +
	`przypisałeś do kategorii: ${categoriesLimited.map((c) => c.slug).join(', ')} i wyszukujesz w nich produkty.`;

export const noSemanticProducts =
	commonPartStartMessage + 'nowe produkty są regularnie dodawane i warto zajrzeć tu za jakiś czas.';

export const foundFirstItems = commonPartStartMessage + 'trwa analiza pierwszych wyników.';

export const noProducts = commonPartStartMessage + 'obecnie nie ma ofert dla pasujących produktów';

export const completingProducts =
	commonPartStartMessage + 'znalazłeś pasujące produkty i za chwilę je wyświetlisz';

export const searchingProductsMessage =
	commonPartStartMessage + 'kompletujesz teraz produkty w tej kategorii';

export const preEmbeddingPrompt = (query: string, availableKeys: string) => `
		Jesteś analitykiem wyszukiwarki e-commerce. Przeanalizuj zapytanie użytkownika: "${query}"
		Dostępne parametry techniczne to: [${availableKeys}]
		
		Zwróć obiekt JSON z 3 polami:
		
		1. "filters": Tablica twardych warunków (np. minimum, maximum, dokładna wartość). 
		   Format: { "key": "nazwa_klucza", "operator": ">=, <=, =, >, <", "value": "wartość" }.
		   Pomiń, jeśli brak takich warunków.
		
		2. "fts_query": Zbiór słów kluczowych do wyszukiwania pełnotekstowego (tylko twarde cechy, bez potocznych słów).
		
		3. "vector_query": Tekst imitujący naturalny opis katalogowy produktu, który idealnie pasuje do zapytania. 
		   WAŻNE: Dopasuj ton do intencji! Jeśli użytkownik szuka czegoś taniego, zacznij np. od "Jest to świetny budżetowy produkt...". Jeśli szuka sprzętu pro, napisz "Zaawansowany sprzęt dla profesjonalistów...". Opisz wymagane cechy w formie płynnych zdań.
	`;
// ANSWER
export const greetingAnswer =
	'C: Jestem  generatorem statusów. ' +
	'O: Chcę stworzyć krótki tekst w formie statusu. ' +
	'S: W stylu profesjonalnym ale nie korporacyjnym. ' +
	'A: Ludzie od 16 roku życia. ' +
	'T: W tonie uprzejmym. ' +
	'R: Przywitaj się z użytkownikiem i spytaj się co dzisiaj szuka w tym samym języku co pytanie użytkownika';

export const endSession =
	'C: Jestem  generatorem statusów. ' +
	'O: Chcę stworzyć krótki tekst w formie statusu. ' +
	'S: W stylu profesjonalnym ale nie korporacyjnym. ' +
	'A: Ludzie od 16 roku życia. ' +
	'T: W tonie uprzejmym. ' +
	'R: Odpowiedz użytkownikowi w tym samym języku co pytanie użytkownika:' +
	'- Wiadomość pozytywna: że się cieszysz i zapraszasz ponownie;' +
	'- Wiadomość negatywna/neutralna: że ci przykro  i postarasz się następnym razem pomóc.';

export const helpersText = (productText: string, hisContent: string) =>
	commonAnswwerStart +
	'Odpowiedz użytkownikowi do 3-4 zdań.' +
	`Opieraj wiedzę TYLKO I WYŁĄCZNIE w oparciu o tą listę produktów: ${productText} i historie jeśli istnieje: ${hisContent}.` +
	'W polu metadata każdego produktu znajdują się parametry techniczne. Nazwy pól mogą być niespójne (np. różne nazwy' +
	' dla wymiarów). Zawsze analizuj te dane kontekstowo, zwracaj uwagę na jednostki (cm, mm, cale), zastosowanie itp ' +
	'i dopasowuj je do wymagań użytkownika (np. wymiary wnęki, moc, pojemność). Jeśli dane są niejasne, uprzedź o tym' +
	' użytkownika. Jeśli danych brak lub są niewystarczające, poinformuj o braku ofert w bazie bez sugerowania ' +
	'zewnętrznych sklepów czy wyszukiwarek. .W razie sprzeczności między opisem a metadata, priorytetem są dane ' +
	'techniczne z metadata.';

export const SEARCH_NEWquestionWhichChangedCategory =
	commonAnswwerStart +
	'Poinformuj dla użytkownika:' +
	'- do 1-2 zdań' +
	'- że wykryto zmianę kategorii.' +
	'- poproś, aby kliknął "Nowy Czat" dla zachowania porządku w historii.' +
	'- za argumentuj to łatwiejszym przeglądaniem rozmów w przyszłości';

export const SEARCH_MOREnewProducts = (newProducts: string, hisContent: string) =>
	commonAnswwerStart +
	'Odpowiedź użytkownikowi:' +
	'- do 3-4 zdań' +
	`- bazując jedynie na danych produktów: ${newProducts}, historii rozmowy, jeśli istnieje: ${hisContent} i jego pytaniu`;

export const SEARCH_MOREnoNewProducts = (products: string, hisContent: string) =>
	commonAnswwerStart +
	'Wyjaśnij użytkownikowi:' +
	'- do 6 zdań' +
	'- dlaczego te modele nadal są najlepsze i dlaczego dalsze szukanie to kompromis jakościowy' +
	`- bazuj tylko na danych produktów: ${products} i historii, jeśli istnieje: ${hisContent}.`;

export const answerForQuestionAboutExistedProducts = (hisContent: string) =>
	commonAnswwerStart +
	'Odpowiedź użytkownikowi:' +
	'- do 3-4 zdań;' +
	`- w oparciu o historie: ${hisContent}` +
	'- jeśli dane w historii są niedostępne (puste), wyjaśnij uprzejmie, że ze względu na limity aktualnego planu' +
	' subskrypcyjnego, Twoja pamięć o poprzednich produktach została odświeżona i poproś o ponowne podanie czego ' +
	'aktualnie szuka';
