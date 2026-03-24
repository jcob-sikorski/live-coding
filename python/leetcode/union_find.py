class UnionFind:
    def __init__(self, size: int):
        if size <= 0:
            raise ValueError("Size must be greater than 0")
        
        self.size = size
        # num_components tracks the total number of disjoint sets
        self.num_components = size
        
        # 'parent[i]' points to the parent of i, if parent[i] == i then i is a root node
        self.parent = list(range(size))
        
        # 'sz[i]' tracks the number of elements in the component rooted at i
        self.sz = [1] * size

    def find(self, p: int) -> int:
        """Finds which component 'p' belongs to, takes amortized constant time."""
        root = p
        while root != self.parent[root]:
            root = self.parent[root]

        # Path Compression: make every node on the path point to the root.
        # This is what gives us the amortized constant time complexity.
        while p != root:
            next_node = self.parent[p]
            self.parent[p] = root
            p = next_node

        return root

    def connected(self, p: int, q: int) -> bool:
        """Returns whether p and q are in the same component."""
        return self.find(p) == self.find(q)

    def component_size(self, p: int) -> int:
        """Returns the size of the component 'p' belongs to."""
        return self.sz[self.find(p)]

    def unify(self, p: int, q: int) -> bool:
        """Unifies the components containing p and q."""
        root1 = self.find(p)
        root2 = self.find(q)

        # These elements are already in the same group
        if root1 == root2:
            return False

        # Merge smaller component into the larger one (Union by Size)
        if self.sz[root1] < self.sz[root2]:
            self.sz[root2] += self.sz[root1]
            self.parent[root1] = root2
        else:
            self.sz[root1] += self.sz[root2]
            self.parent[root2] = root1

        # Since we merged two components, the total count decreases
        self.num_components -= 1
        return True