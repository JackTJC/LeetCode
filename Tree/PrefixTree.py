# 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
#
# 示例:
#
# Trie trie = new Trie();
#
# trie.insert("apple");
# trie.search("apple");   // 返回 true
# trie.search("app");     // 返回 false
# trie.startsWith("app"); // 返回 true
# trie.insert("app");
# trie.search("app");     // 返回 true
# 说明:
#
# 你可以假设所有的输入都是由小写字母 a-z 构成的。
# 保证所有输入均为非空字符串。
#
# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

class Trie:

    def __init__(self, char=""):
        self.char = char
        self.isWord = False
        self.children = [None] * 26

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        cur = self
        for c in word:
            nextNode = cur.children[ord(c) - ord('a')]
            if nextNode:
                cur = nextNode
            else:
                cur.children[ord(c) - ord('a')] = Trie(c)
                cur = cur.children[ord(c) - ord('a')]
        cur.isWord = True

    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        cur = self
        for c in word:
            nextNode = cur.children[ord(c) - ord('a')]
            if nextNode is None:
                return False
            cur = nextNode
        return cur.isWord

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        cur = self
        for c in prefix:
            nextNode = cur.children[ord(c) - ord('a')]
            if nextNode is None:
                return False
            cur = nextNode
        return True
