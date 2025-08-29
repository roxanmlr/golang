# Explications
## nothing.go
Un programme Go peut se limiter à un seul fichier source.
Ce fichier doit :

* avoir l’extension **`.go`**,
* commencer par la déclaration du paquet principal :

```go
package main
```

* contenir une fonction `main`, point d’entrée obligatoire du programme :

```go
func main() {
    // instructions
}
```

Un exemple minimal est le fichier **`nothing.go`** : il constitue un programme complet en Go, même s’il n’exécute aucune action.

## print.go

### 1. Le package `main`

En Go, tout fichier source commence par la déclaration d’un **package**.
Un **package** est un regroupement logique de code : fonctions, types, constantes…

* Lorsqu’on écrit une bibliothèque réutilisable, on choisit un nom de package (ex. `math`, `fmt`).
* Lorsqu’on écrit un programme exécutable, on utilise toujours le **package `main`**.

N.B: Le package `main` est spécial : il doit contenir la fonction `main`, point d’entrée de l’application.
C’est à partir de cette fonction que l’exécution du programme démarre.

Exemple :

```go
package main
```

### 2. Importer la bibliothèque standard

Go est livré avec une riche **bibliothèque standard** (standard library) : elle fournit déjà des outils pour gérer les entrées/sorties, les fichiers, le réseau, les chaînes de caractères, etc.

On utilise le mot-clé **`import`** pour indiquer quelles parties de cette bibliothèque on veut utiliser.

Exemple :

```go
import "os"
```

Ici, on importe le package **`os`** qui permet de travailler avec le système d’exploitation (fichiers, variables d’environnement, sorties, etc.).

### 3. La fonction `main` et la sortie standard

La fonction `main` est obligatoire dans un programme exécutable :

```go
func main() {
    // instructions
}
```

Dans l’exemple suivant, on écrit un caractère directement sur la **sortie standard** :

```go
package main

import "os"

func main() {
    os.Stdout.Write([]byte{'c'})
}
```

### Qu’est-ce que la sortie standard ?

* La **sortie standard** (souvent notée **stdout**) est un flux (stream) dans lequel un programme écrit ses résultats.
* Par défaut, ce flux est relié à l’écran (le terminal où vous lancez le programme).

Ici, `os.Stdout.Write` écrit le caractère `'c'` dans ce flux, donc le programme affiche simplement :

```
c
```

---

En résumé :

* `package main` → indique qu’on écrit un programme exécutable.
* `func main()` → point d’entrée du programme.
* `import` → permet d’utiliser la bibliothèque standard.
* `os.Stdout.Write` → écrit sur la sortie standard (l’écran).
Très bien. Voici un cours structuré autour du programme `print_easy.go`, avec explications progressives :

---

## print\_easy.go

### 1. Importation des packages

```go
import (
	"fmt"
	"os"
)
```

Ici, on importe deux packages de la bibliothèque standard :

* **`fmt`** : fournit des fonctions de formatage et d’affichage (par exemple `Println`, `Printf`).
* **`os`** : permet d’interagir avec le système d’exploitation, en particulier les arguments passés en ligne de commande.

### 2. La fonction `main`

```go
func main() {
	...
}
```

Comme toujours dans un programme exécutable, l’exécution commence dans la fonction `main`.

### 3. Les variables et l’initialisation

```go
var name string = "unknown"
```

Ici, on déclare une variable `name` de type `string`, initialisée à `"unknown"`.

Go propose aussi une syntaxe abrégée avec `:=` qui permet de déclarer et d’initialiser en une seule fois :

```go
age := "unknown"
```

### 4. Les arguments de la ligne de commande (`os.Args`)

* `os.Args` est une **liste** (plus exactement un *slice*) de chaînes de caractères.
* Elle contient tous les arguments donnés au programme lors de son exécution.
* `os.Args[0]` est toujours le nom du programme.
* Les arguments suivants (`os.Args[1]`, `os.Args[2]`, …) sont ceux fournis par l’utilisateur.

Exemple :

```bash
go run print_easy.go Alice 30
```

Alors :

* `os.Args[0]` = `"print_easy.go"`
* `os.Args[1]` = `"Alice"`
* `os.Args[2]` = `"30"`

### 5. Le branchement conditionnel `if`

Le programme vérifie si l’utilisateur a donné assez d’arguments avant d’accéder à `os.Args[1]` ou `os.Args[2]`.

```go
if len(os.Args) >= 2 {
	name = os.Args[1]
}
if len(os.Args) >= 3 {
	age = os.Args[2]
}
```

Cela évite une erreur si les arguments manquent.

* Si au moins 1 argument est présent, on affecte `name`.
* Si au moins 2 arguments sont présents, on affecte `age`.

### 6. L’affichage formaté avec `fmt.Printf`

```go
fmt.Printf("Hello %v, You are %v years old", name, age)
```

`fmt.Printf` permet d’afficher du texte avec des variables insérées à l’intérieur.

* `%v` est un **verbe de formatage** : il affiche une valeur dans sa forme la plus lisible.
* Ici, `%v` est remplacé par la valeur de `name`, puis de `age`.

Exemple d’exécution :

```bash
go run print_easy.go
```

Affiche :

```
Hello unknown, You are unknown years old
```

```bash
go run print_easy.go Alice 30
```

Affiche :

```
Hello Alice, You are 30 years old
```

---

### Résumé

Ce programme introduit :

* les **packages** `fmt` (affichage formaté) et `os` (accès aux arguments),
* la **liste** `os.Args` contenant les arguments de la ligne de commande,
* le **branchement conditionnel** `if`,
* la déclaration et l’initialisation de variables,
* l’usage de `fmt.Printf` pour produire une sortie personnalisée.
Parfait. Voici le cours structuré autour du programme **`print_n_times.go`** :

---

## print\_n\_times.go

### 1. Importation des packages

```go
import (
	"fmt"
	"os"
	"strconv"
)
```

* **`fmt`** : affichage formaté.
* **`os`** : récupération des arguments de la ligne de commande.
* **`strconv`** : conversion entre chaînes de caractères et types numériques.

### 2. Vérification des arguments

```go
if len(os.Args) != 3 {
	fmt.Printf("Usage:%v string_to_print times\n", os.Args[0])
	// os.Args[0] est le nom du programme
	return
}
```

Ici, on exige exactement deux arguments après le nom du programme :

1. `string_to_print` : la chaîne à afficher,
2. `times` : combien de fois l’afficher.

* Si le nombre d’arguments n’est pas correct, le programme affiche un **message d’utilisation** (*usage*) et s’arrête.
* Exemple :

```bash
go run print_n_times.go
```

Affiche :

```
Usage:print_n_times.go string_to_print times
```

### 3. Conversion de chaîne en entier

```go
n, _ := strconv.Atoi(os.Args[2])
```

* `os.Args[2]` est une chaîne (par exemple `"5"`).
* `strconv.Atoi` convertit cette chaîne en un entier.
* Le résultat est stocké dans `n`.
* Le second retour `_` est l’erreur éventuelle (ici on l’ignore).

### 4. La boucle `for`

```go
for i := 0; i < n; i++ {
	fmt.Println(i+1, ": ", os.Args[1])
}
```

En Go, la boucle `for` est la seule structure de répétition. Ici :

* `i := 0` initialise le compteur,
* `i < n` indique la condition d’arrêt,
* `i++` incrémente `i` à chaque tour.

À chaque itération, on affiche :

* le numéro de l’itération (`i+1`),
* suivi de la chaîne donnée par l’utilisateur (`os.Args[1]`).

### 5. Exemple d’exécution

Commande :

```bash
go run print_n_times.go Hello 3
```

Affiche :

```
1 :  Hello
2 :  Hello
3 :  Hello
```

### Résumé

Ce programme introduit :

* la vérification du nombre d’arguments pour guider l’utilisateur,
* la conversion de chaînes en entiers avec `strconv.Atoi`,
* la boucle `for`, qui permet de répéter des instructions un nombre précis de fois.

Très bien. Voici le cours correspondant à ce programme, en me concentrant uniquement sur les **nouvelles notions** introduites :

---

## table.go

### La boucle `for … range` sur un entier

Dans Go, `range` sert généralement à parcourir une **liste** (slice, tableau, chaîne, map).
Depuis Go 1.22, il est aussi possible de l’utiliser directement sur un entier `n` :

```go
for i := range n {
    // i prend successivement les valeurs 0, 1, ..., n-1
}
```

C’est équivalent à écrire :

```go
for i := 0; i < n; i++ {
    ...
}
```

Dans ce programme, `for mul := range n` parcourt toutes les valeurs de `0` à `n-1`.
On ajoute `+1` pour commencer les tables à 1.

---

### Les fonctions définies par l’utilisateur

En Go, on peut définir ses propres fonctions pour organiser le code et éviter les répétitions.
La syntaxe est :

```go
func nom_fonction(param1 type1, param2 type2, ...) typeRetour {
    // instructions
    return valeur
}
```

Ici, on définit une fonction `print_table` :

```go
func print_table(mul int, max int) {
	for i := range max + 1 {
		fmt.Printf("%v x %v = %v\n", mul, i, mul*i)
	}
}
```

* Elle reçoit deux paramètres :

  * `mul` → le nombre dont on veut afficher la table,
  * `max` → jusqu’où on veut calculer la table de multiplication.
* Elle n’a pas de valeur de retour, elle se contente d’afficher.

---

### Exemple d’exécution

Commande :

```bash
go run table.go 5 3
```

Résultat :

```
Table de 1
1 x 0 = 0
1 x 1 = 1
1 x 2 = 2
1 x 3 = 3
1 x 4 = 4
1 x 5 = 5

Table de 2
2 x 0 = 0
2 x 1 = 2
2 x 2 = 4
2 x 3 = 6
2 x 4 = 8
2 x 5 = 10

Table de 3
3 x 0 = 0
3 x 1 = 3
3 x 2 = 6
3 x 3 = 9
3 x 4 = 12
3 x 5 = 15
```

---

### Résumé

Dans ce programme, on découvre :

* l’utilisation de `for … range` directement sur un entier,
* la définition de fonctions utilisateurs avec `func`,
* le passage de paramètres aux fonctions pour réutiliser du code.
