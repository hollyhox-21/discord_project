window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  let urls = [
      {url: '/docs/router.swagger.json', name: "doc1"},
      {url: '/docs/pet.swagger.json', name: "doc2"},
  ];

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: urls,
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
