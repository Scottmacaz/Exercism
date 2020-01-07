
def _get_column(matrix, column):
    c = []
    for row in matrix:
        c.append(row[column])
    return c

def _is_matrix_regular(matrix):
    if not matrix:
        return True
    
    matrixlen = len(matrix[0])
    for row in matrix:
        if matrixlen != len(row):
            return False
    return True

def saddle_points(matrix):
    x=0
    if not _is_matrix_regular(matrix):
        raise ValueError("Matrix is not regular.")
    saddlepoints = []
    for row in matrix:
        for y in range(len(row)):
            column = _get_column(matrix, y)
            if row[y] == min(column)  and row[y] == max(row) :
                saddlepoints.append({"row": x+1, "column": y+1})
        x += 1
    return saddlepoints
