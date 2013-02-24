package handler

import (
    "fmt"
    "net/http"
    "./data"
    "strconv"
)

func RegisterRoutesProject() {
    RegisterRoute("^/([^/]+)$", func(w http.ResponseWriter, r *http.Request, captures []string) {
        projectName := captures[0]
        fmt.Println("project", projectName)

        params := make(map[string]interface{})
        params["topProject"] = data.GetProject("manage")
        params["project"] = data.GetProject(projectName)
        params["wiki"] = data.GetWiki(projectName, "top")
        params["newestTickets"] = data.GetNewestTickets(projectName, 10)
        params["states"] = data.GetStates(projectName, false)

        TmplProject(w, "project", params)
    })

    RegisterRoute("^/([^/]+)/list", func(w http.ResponseWriter, r *http.Request, captures []string) {
        projectName := captures[0]
        fmt.Println("project list", projectName)

        params := make(map[string]interface{})
        params["topProject"] = data.GetProject("manage")
        params["project"] = data.GetProject(projectName)
        params["newestTickets"] = data.GetNewestTickets(projectName, 10)
        params["states"] = data.GetStates(projectName, false)

        notClosedStates := data.GetStates(projectName, true)
        ticketsByStatus := []interface{}{}
        for _, s := range notClosedStates {
            statusResult := map[string]interface{}{}
            statusResult["State"] = s
            statusResult["SearchResult"] = data.GetTicketsByStatus(projectName, s.Name)
            ticketsByStatus = append(ticketsByStatus, statusResult)
        }
        params["notClosedStates"] = notClosedStates
        params["ticketsByStatus"] = ticketsByStatus
        params["elementTypes"] = data.GetElementTypes(projectName)
        TmplProject(w, "project_list", params)
    })

    RegisterRoute("^/([^/]+)/setting_file/([^/]+)", func(w http.ResponseWriter, r *http.Request, captures []string) {
        projectName := captures[0]
        name := captures[1]

        file := data.GetSettingFile(projectName, name)

        w.Header().Set("Content-Type", file.MimeType)
        w.Header().Set("Content-Length", strconv.Itoa(file.Size))
        w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", 60 * 24 * 1))

        w.Write([]byte(file.Content))
    })

}

/* vim: set ts=4 sw=4 sts=4 expandtab fenc=utf-8: */
