

def recebe(valor):
    base = valor
    while (valor > 1):
        base = (valor * (valor -1))
    return base 


print(recebe(10))