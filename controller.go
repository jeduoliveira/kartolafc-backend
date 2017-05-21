package kartolafc

import (
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"net/http"
	"github.com/pressly/chi"
	"strconv"
	"github.com/pressly/chi/render"
	"encoding/json"
)

func GetHome(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	var out map[string]string
	json.Unmarshal([]byte("{ \"status\":\"Birll\"}"), &out)
	render.JSON(response, request, out)
}

func GetStatus(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	render.JSON(response, request, CacheStatus)
}

func GetTimes(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	timePesquisado := chi.URLParam(request, "q")
	times := api.Times{}
	times.Pesquisa = timePesquisado
	times.GetTimes()

	response.Write(JsonBuild(times))
}

func GetTime(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	idString := chi.URLParam(request, "id")
	time := api.TimeCompleto{}

	id, err := strconv.Atoi(idString)

	if err != nil {
		response.Write([]byte("{\"status\": \"error\", \"message\": \"id tem que ser um numero\"}"))
	} else {
		time.TimeCompleto.TimeId = id
		time.GetTime()
		render.JSON(response, request, time)
	}
}

func GetMercado(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	render.JSON(response, request, CacheKartolaAtletas)
}

func GetDestaques(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	render.JSON(response, request, CacheDestaques)
}

func GetLigas(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	nome := chi.URLParam(request, "q")
	pesquisa := api.PesquisaLigas{}
	pesquisa.GetPesquisaLigas(nome)
	render.JSON(response, request, pesquisa)
}

func GetLiga(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	var page int
	pageString := chi.URLParam(request, "page")
	if pageString != "" {
		pageOne , err := strconv.Atoi(pageString)
		if err != nil || pageOne < 1 || pageOne > 5 {
			var out map[string]string
			json.Unmarshal([]byte("{ \"status\":\"error\", \"message\":\"id informado nao é numerico ou menor que 1 ou maior que 5\"}"), &out)
			render.JSON(response, request, out)
			return;
		}
		page = pageOne
	}
	slug := chi.URLParam(request, "id")
	liga := api.Liga{}
	liga.Liga.Slug = slug
	liga.GetLiga(page)
	render.JSON(response, request, liga)
}

func GetPontuados(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)

	render.JSON(response, request, CachePontuados)
}

func responseDefault(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}