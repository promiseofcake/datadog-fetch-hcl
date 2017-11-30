resource "datadog_timeboard" "search_service" {
  title = "Search Service"
  description = "Search Service"
  read_only = false

  template_variable {
    name    = "ENV"
    prefix  = "environment"
    default = "environment:prod"
  }

  graph {
    title = "Elasticsearch QPS"
    autoscale = true
    viz = "timeseries"

    request {
      q = "sum:search_api.related_images.search.count{$ENV}.as_count()"
      type = "line"
    }
  }

  graph {
    title = "Related Image RPS to Store Search RPS (hitpass)"
    autoscale = true
    viz = "timeseries"

    request {
      q = "sum:search_api.related_images.search.count{$ENV}.as_rate()"
      type = "line"
    }

    request {
      q = "sum:search.api.http.get.2_0.search.images.id.related.count{$ENV}.as_rate()"
      type = "line"
    }
  }
}
