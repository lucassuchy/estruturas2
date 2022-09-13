from ProdStruc import ProdStruc

class Node:
    def __init__(self, data, proximo=None):
        self.data = data
        self.proximo = proximo
    def setData(self,data):
        set.data = data
    def getData(self):
        return self.data
    def setProximo(self, proximo):
        self.proximo = proximo
    def getProximo(self):
        return self.proximo
    
class Lista:
    def __init__(self):
        self.head = None
    def retornaLista(self):
        temp = self.head
        while(temp):
            print(temp.data)
            temp = temp.proximo
    def insertElementoInicio(self, data):
        # preciso ver pq isso não ta funcionando
        if self.head == None:
            no = Node(data)
            self.head = no
        else:
            no = Node(data)
            no.setProximo = self.head
            self.head =  no
    def insertElementoFinal(self, data):
        no = Node(data)
        temp = self.head
        while(temp.proximo != None):
            temp = temp.proximo
        temp.proximo = no


class Controller:
# Vou usar esse metodo pra alimentar a lista
    def __init__(self,prod):
        self.prod = prod

    def __str__(self):
        return f"Nome: {self.nome} \nTamanho: {self.tamanho} \nPreço: {self.preco} \nPeso: {self.peso}\nDescrição:{self.descricao}"

    def recebeProduto(self, nome, tamanho, preco, peso, descricao):
        self.nome = nome
        self.tamanho = tamanho
        self.preco = preco
        self.peso = peso
        self.descricao = descricao
        return self

    def appendaLista(self):
        lista = Lista()
        lista.insertElementoFinal(self)




