class TrieNode:
    def __init__(self, val: str):
        self.val = val
        self.leafs = {}
        self.complete_word = False

    def add(self, val: str):
        if val not in self.leafs:
            self.leafs[val] = TrieNode(self.val + val)

        return self.leafs[val]

    def leafs(self):
        return self.leafs

    def has_leaf(self, val: str) -> bool:
        return val in self.leafs

    def get_leaf(self, val: str) -> 'TrieNode':
        return self.leafs[val]

    def mark_as_word(self):
        self.complete_word = True

    def is_word(self):
        return self.complete_word


class Trie:
    def __init__(self):
        self.root = TrieNode('')

    def insert(self, word: str) -> None:
        current_node = self.root
        for letter in word:
            current_node = current_node.add(letter)

        current_node.mark_as_word()

    def search(self, word: str) -> bool:
        current_node = self.root
        for letter in word:
            if current_node.has_leaf(letter):
                current_node = current_node.get_leaf(letter)
            else:
                return False

        return current_node.is_word()

    def startsWith(self, prefix: str) -> bool:
        current_node = self.root
        for letter in prefix:
            if current_node.has_leaf(letter):
                current_node = current_node.get_leaf(letter)
            else:
                return False

        return True

    def print(self) -> None:
        print(self.root)


# Your Trie object will be instantiated and called as such:
obj = Trie()
obj.insert('apple')
obj.insert('app')
obj.insert('apps')
obj.insert('axe')
print(obj.search('axe'))
print(obj.search('ap'))
print(obj.search('app'))
print(obj.search('apps'))
print(obj.search('a'))
print(obj.startsWith('ap'))
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)


test = {}
