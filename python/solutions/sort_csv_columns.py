def sort_csv_columns(csv_file_content: str) -> str:
    result = ''
    rows = csv_file_content.split("\n")
    order = get_order(rows[0])
    for row in rows:
        row = row.split(";")
        new_row = [row[index] for index in order]
        result += ";".join(new_row) + "\n"

    return result.rstrip()


def get_order(row) -> tuple[int, ...]:
    ordered_row = sorted(
        [(name, index) for index, name in enumerate(row.split(";"))],
        key=lambda name_and_index: name_and_index[0].lower()
    )
    return tuple(index for name, index in ordered_row)

if __name__ == '__main__':
    data = (
        "myjinxin2015;raulbc777;smile67;Dentzil;SteffenVogel_79\n"
        "17945;10091;10088;3907;10132\n"
        "2;12;13;48;11"
    )
    print(sort_csv_columns(data))