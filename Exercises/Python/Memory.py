students = [
    dict(id=0, credits=dict(math=9, physics=6, history=7)),
    dict(id=1, credits=dict(math=6, physics=7, latin=10)),
    dict(id=2, credits=dict(history=8, physics=9, chemistry=10)),
    dict(id=3, credits=dict(math=5, physics=5, geography=7))
]

## Create a tuple of the sum of credits and students
def decorate(student):
    return (sum(student['credits'].values()), student)


def undecorate(decorated_student):
    # discard sum of credits, return original student dict
    return decorated_student[1]


students = sorted(map(decorate, students), reverse=True)
students = list(map(undecorate, students))

print(decorate(students[0]))

print(decorate, students)


## filter function 

## Removes all repeating none elements
test = [2, 5, 8, 0, 0, 1, 0]
list(filter(None, test))



## Comprehensions - 3 different types, list, dict and set
## List example 1
squares = []

for n in range(0, 10):
    squares.append(n**2)

list(squares)

## List example 2 
squares = map(lambda  n: n**2, range(10))
list(squares)


## Filtering a comprehension

from math import sqrt

## will generate all possible pairs

max = 10
legs = [(a, b, sqrt(a**2 + b**2))
    for a in range(1, max) for b in range(a, max)]

# this will filter out all non pythagorean triples
legs = list(
    filter(lambda triple: triple[2].is_integer(), legs))
print(legs)


## Dict comprehensions

from string import ascii_lowercase
lettermap = dict((c, k) for k, c in enumerate(ascii_lowercase, 1))
print(lettermap)

