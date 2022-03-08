INSERT INTO content
 ( id, created_at, updated_at, deleted_at,
   url, slug, title, format,
   summary,
   description,
   archetype, section,
   content,
   language,
   front_matter, path, layout
 )
 VALUES
 ( "1","0","0","",
   "/test","test","Testing title","tmpl",
   "Summary test",
   "Described test",
   "_default","",
   "This content has a custom variable: '{{ .CustomVariable }}'.",
   "en_CA",
   "","",""
 )
;
