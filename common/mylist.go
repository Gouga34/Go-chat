package common

//MyList permet d'avoir toutes des fonctions sur un slice
type MyList struct {
	array []interface{}
}

//Cut enlève du slice tous les éléments compris entre i et j
func (list *MyList) Cut(i, j int) {
	list.array = append(list.array[:i], list.array[j:]...)
}

//Copy copie la liste listToCopy dans list
func (list *MyList) Copy(listToCopy MyList) {
	list.array = make([]interface{}, len(listToCopy.array))
	copy(list.array, listToCopy.array)
}

//Delete supprime l'élément à la position i de list
func (list *MyList) Delete(i int) {
	list.array = append(list.array[:i], list.array[i+1:]...)
}

//Pop enlève le dernier élément ajouté et le retourne
func (list *MyList) Pop() interface{} {
	var x interface{}
	x, list.array = list.array[len(list.array)-1], list.array[:len(list.array)-1]
	return x
}

//Push ajoute newElement à la fin
func (list *MyList) Push(newElement interface{}) {
	list.array = append(list.array, newElement)
}
