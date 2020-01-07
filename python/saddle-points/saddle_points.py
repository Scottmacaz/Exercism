
def _get_column(matrix, column):
    c = []
    for row in matrix:
        c.append(row[column])
    return c

def _is_matrix_regular(matrix):
    matrixlen = len(matrix[0])
    for row in matrix:
        if matrixlen != len(row):
            return False
    return True

def saddle_points(matrix):
    saddlepoints = []
    if not matrix:
        return saddlepoints
    if not _is_matrix_regular(matrix):
        raise ValueError("Matrix is not regular.")
    
    for x, row in enumerate(matrix, start=1):
        for y in range(len(row)):
            column = _get_column(matrix, y)
            if row[y] == min(column)  and row[y] == max(row) :
                saddlepoints.append({"row": x, "column": y+1})
    return saddlepoints
