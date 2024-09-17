package helper

// GeneratePaginationMeta membuat map untuk meta informasi pagination
func GeneratePaginationMeta(totalCount, currentPage, totalPages, nextPage, prevPage int) map[string]interface{} {
    if currentPage <= 1 {
        prevPage = 0
    }
    if currentPage >= totalPages {
        nextPage = 0
    }

    return map[string]interface{}{
        "total_records": totalCount,
        "current_page":  currentPage,
        "total_pages":   totalPages,
        "next_page":     nextPage,
        "prev_page":     prevPage,
    }
}
