class Matrix:
    def __init__(self, matrix_string):
        rows = matrix_string.split("\n")
        self._matrix = []
        for row in rows:
            self._matrix.append([int(x) for x in row.split(" ")])

    def row(self, index):
        return self._matrix[index-1]

    def column(self, index):
        col = []
        for row in self._matrix:
            col.append(row[index-1])
        return col
