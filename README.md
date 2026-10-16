# Go-Reloaded

## Table des matières
1. [Description du projet](#description-du-projet)
2. [Fonctionnalités et Règles de Transformation](#fonctionnalités-et-règles-de-transformation)
3. [Architecture Technique](#architecture-technique)
4. [Arborescence du Projet](#arborescence-du-projet)
5. [Installation et Prérequis](#installation-et-prérequis)
6. [Utilisation](#utilisation)
7. [Exemple Complet de Test](#exemple-complet-de-test)
8. [Méthodologie de Développement (Git)](#méthodologie-de-développement-git)

---

## Description du projet
**Go-Reloaded** est un outil de traitement de texte intelligent développé en langage Go. Son objectif principal est de recevoir un fichier texte brut contenant diverses balises de formatage, des nombres à convertir, ou des erreurs de ponctuation/grammaire, et de générer un fichier de sortie parfaitement corrigé et mis en page. 

Le programme respecte scrupuleusement la structure originale du document (en préservant les sauts de ligne et les lignes vides) tout en effectuant un nettoyage syntaxique rigoureux.

---

## Fonctionnalités et Règles de Transformation

Le programme analyse le texte mot par mot et ligne par ligne pour appliquer les règles suivantes :

1. **Conversions de bases numériques :**
   - `(hex)` : Convertit le mot précédent, exprimé en hexadécimal, en sa valeur décimale.
   - `(bin)` : Convertit le mot précédent, exprimé en binaire, en sa valeur décimale.

2. **Modificateurs de casse :**
   - `(up)` : Met le mot précédent entièrement en majuscules.
   - `(low)` : Met le mot précédent entièrement en minuscules.
   - `(cap)` : Met la première lettre du mot précédent en majuscule.
   - `(up, n)`, `(low, n)`, `(cap, n)` : Applique la modification de casse correspondante sur les $n$ mots précédents.

3. **Correction des articles indéfinis :**
   - Transforme automatiquement l'article `a` en `an` si le mot qui le suit commence par une voyelle (`a, e, i, o, u`) ou par un `h`.

4. **Formatage de la ponctuation :**
   - Recolle les signes de ponctuation (`.`, `,`, `!`, `?`, `:`, `;`) directement au mot précédent et s'assure qu'un espace unique suit la ponctuation.
   - Traite les groupes de ponctuation spécifiques comme les points de suspension (`...`) ou les points d'interrogation/exclamation combinés (`!?`).
   - Corrige l'espacement autour des guillemets simples (`'`) pour qu'ils soient collés au texte qu'ils encadrent, tout en gérant leur parité.

5. **Préservation de la mise en page :**
   - Grâce à un traitement ligne par ligne, les retours à la ligne et les lignes vides du fichier d'origine sont intégralement conservés dans le fichier de sortie.

---

## Architecture Technique

Le code source est divisé en deux modules principaux pour assurer une séparation claire des responsabilités :
- **`main.go`** : Gère la vérification des arguments de la ligne de commande, l'ouverture, la lecture du fichier d'entrée et l'écriture sécurisée dans le fichier de sortie.
- **`modif.go`** : Contient le cœur de la logique algorithmique (`Process`, `formatPunctuation`, `capitalize`) et utilise le lecteur de flux (`bufio.Scanner`) pour itérer proprement sur chaque ligne du texte.

---

## Arborescence du Projet

```text
go-reloaded/
│
├── entree.txt          # Fichier source contenant le texte brut à modifier
├── sortie.txt          # Fichier généré automatiquement avec le texte corrigé
├── go.mod              # Fichier de configuration du module Go
├── main.go             # Point d'entrée de l'application (I/O et gestion des fichiers)
├── modif.go            # Moteur de transformation et de formatage textuel
└── README.md           # Documentation complète du projet