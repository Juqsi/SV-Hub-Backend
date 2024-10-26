package main

import (
	"HexMaster/api/routes"
	"HexMaster/weaviate"
)

func main() {

	weaviate.DeleteSchema()

	texts := []string{
		"Das Mittelalter war eine Epoche in der europäischen Geschichte, die etwa von 500 bis 1500 n. Chr. andauerte. Es war geprägt von Feudalismus und der Macht der Kirche.",

		"Im Mittelalter war die Religion ein zentraler Bestandteil des Lebens. Klöster spielten eine wichtige Rolle bei der Bildung und Aufbewahrung von Wissen.",

		"Die Architektur im Mittelalter war beeindruckend. Gotische Kathedralen wie der Kölner Dom zeigen die Meisterwerke der Baukunst dieser Zeit.",

		"Ritter waren bedeutende Figuren des Mittelalters. Sie folgten einem Ehrenkodex und kämpften oft in Turnieren und Kriegen für ihre Lehnsherren.",

		"Das Alltagsleben im Mittelalter war hart. Bauern arbeiteten auf Feldern und hatten oft wenig Rechte, während Adlige in Burgen lebten und das Land beherrschten.",

		"Die Pest, auch Schwarzer Tod genannt, wütete im 14. Jahrhundert und führte zum Tod von Millionen von Menschen in Europa.",

		"Handel und Märkte entwickelten sich im Mittelalter und trugen zum Wachstum von Städten bei. Städte wie Venedig wurden zu wichtigen Handelszentren.",

		"Die Kunst im Mittelalter war stark religiös geprägt. Ikonen und Wandmalereien schmückten Kirchen und Klöster und zeigten biblische Szenen.",

		"Burgen und Festungen wurden im Mittelalter gebaut, um die Bevölkerung zu schützen. Diese Bauwerke sind heute noch in vielen Teilen Europas zu finden.",

		"Das Lehnswesen war ein zentraler Bestandteil der mittelalterlichen Gesellschaft. Lehnsherren vergaben Land an ihre Vasallen im Austausch für Treue und militärische Unterstützung.",

		"Kreuzzüge wurden im Mittelalter von der Kirche initiiert, um das Heilige Land zurückzuerobern. Sie führten zu langen und oft blutigen Auseinandersetzungen.",

		"Bildung war im Mittelalter vor allem auf die Kleriker beschränkt. Universitäten entstanden jedoch gegen Ende der Epoche, etwa in Paris und Bologna.",

		"Musik spielte eine bedeutende Rolle im Mittelalter. Sie wurde oft von Mönchen in Klöstern gesungen und diente der religiösen Andacht.",

		"Im späten Mittelalter entstanden Gilden, die die Interessen von Handwerkern und Kaufleuten schützten. Diese Organisationen waren entscheidend für die städtische Wirtschaft.",

		"Das Ende des Mittelalters wurde durch die Renaissance eingeläutet, eine Epoche, die das Interesse an Kunst und Wissenschaft neu entfachte.",
	}

	// Insert texts into Weaviate and calculate vectors
	weaviate.InsertData(texts)
	routes.SetupRoutes()
	//fmt.Println(llama.DoRequestWithVectors("Erzähl mir über das Mittelalter"))
	/*
		resp, _ := llama.DoRequest("Welche Suchbegriffe sind gut für folgende frage für eine Vektor atenbank: '" + "Haben wir bequeme T-Shirts im Sorrtiment" + "'? Gibt bitte die begriffe in Eckigenklammern zurück und darin mit einem Komma getrennt. Bitte schreibe ausschließlich diese syntax und keine andere auch nicht für antowrt text")
		fmt.Println(resp)
		lis := strings.Split(resp[1:len(resp)-1], ",")
		fmt.Println("+++++++++++++++++++")
		for _, li := range lis {
			fmt.Println(weaviate.RetrieveVectors("Document", "text", li))
		}

		// Beispiel für die Verwendung der doRequest-Funktion
		/*prompt := "Wie sehen die Finanzenumschläge der Schülervertretung vom letzten Jahr aus"
		println("Starte Programm")
		response, err := llama.DoRequest(prompt)
		if err != nil {
			fmt.Printf("Fehler: %v\n", err)
			return
		}

		fmt.Printf("Antwort von LLaMA: %s\n", response)
	*/
}
