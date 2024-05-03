package diccionario

import "fmt"

type estadoNodo int

const (
	VACIO = estadoNodo(iota)
	OCUPADO
	BORRADO
)
const LARGOINICIAL int = 13
const PANICO = "La clave no pertenece al diccionario"
const PANICOITER = "El iterador termino de iterar"

type nodoHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoNodo
}

type hashCerrado[K comparable, V any] struct {
	contenido []nodoHash[K, V]
	tamaño    int
	ocupados  int
	borrados  int
}

type iterDiccionarioHash[K comparable, V any] struct {
	hash     *hashCerrado[K, V]
	posicion int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	nuevo := new(hashCerrado[K, V])
	nuevo.contenido = make([]nodoHash[K, V], LARGOINICIAL)
	nuevo.tamaño = LARGOINICIAL
	nuevo.ocupados, nuevo.borrados = 0, 0
	return nuevo
}

// Primitivas del Hash
func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

}
func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {

}
func (hash *hashCerrado[K, V]) Obtener(clave K) V {

}
func (hash *hashCerrado[K, V]) Borrar(clave K) V {

}
func (hash *hashCerrado[K, V]) Cantidad() int {

}
func (hash *hashCerrado[K, V]) Iterar(func(clave K, dato V) bool) {

}
func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

}

// Primitivas del iterador
func (iter *iterDiccionarioHash[K, V]) HaySiguiente() bool {
	return iter.posicion < iter.hash.tamaño && iter.hash.contenido[iter.posicion].estado == OCUPADO

}
func (iter *iterDiccionarioHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(PANICOITER)
	}
	elem := iter.hash.contenido[iter.posicion]
	return elem.clave, elem.dato
}
func (iter *iterDiccionarioHash[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(PANICOITER)
	}
	tabla := iter.hash.contenido
	for ; (iter.posicion < iter.hash.tamaño) && (tabla[iter.posicion].estado != OCUPADO); iter.posicion++ {
	}
}

//funciones auxiliares

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func indiceHash[K comparable](clave K, largo_tabla int) int { //usando la funcion hash me da el indice en que inicialmente le corresponde a mi valor
	convertirABytes(clave)
	//hacer el hash
	//hacemos return hash % largo_tabla
}

// Recibe una cantidad para redimensionar
func redimensionar[K comparable, V any](hash *hashCerrado[K, V], cantidad_nueva int) {
	nuevo_contenido := make([]nodoHash[K, V], cantidad_nueva)
	for _, e := range hash.contenido {
		if e.estado == OCUPADO {
			ind_ini := indiceHash(e.clave, cantidad_nueva)
			indice := ind_ini
			for ; indice < cantidad_nueva && nuevo_contenido[indice].estado == OCUPADO; indice++ {
			}
			if indice == cantidad_nueva {
				indice = 0
				for ; indice < ind_ini && nuevo_contenido[indice].estado == OCUPADO; indice++ {
				}
			}
			nuevo_contenido[indice].clave, nuevo_contenido[indice].dato = e.clave, e.dato
		}
	}
	hash.contenido = nuevo_contenido
	hash.tamaño = cantidad_nueva
	hash.borrados = 0
}
