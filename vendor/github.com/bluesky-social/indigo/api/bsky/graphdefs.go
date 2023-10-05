// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.defs

// GraphDefs_ListItemView is a "listItemView" in the app.bsky.graph.defs schema.
type GraphDefs_ListItemView struct {
	Subject *ActorDefs_ProfileView `json:"subject" cborgen:"subject"`
}

// GraphDefs_ListView is a "listView" in the app.bsky.graph.defs schema.
//
// RECORDTYPE: GraphDefs_ListView
type GraphDefs_ListView struct {
	LexiconTypeID     string                     `json:"$type,const=app.bsky.graph.defs#listView" cborgen:"$type,const=app.bsky.graph.defs#listView"`
	Avatar            *string                    `json:"avatar,omitempty" cborgen:"avatar,omitempty"`
	Cid               string                     `json:"cid" cborgen:"cid"`
	Creator           *ActorDefs_ProfileView     `json:"creator" cborgen:"creator"`
	Description       *string                    `json:"description,omitempty" cborgen:"description,omitempty"`
	DescriptionFacets []*RichtextFacet           `json:"descriptionFacets,omitempty" cborgen:"descriptionFacets,omitempty"`
	IndexedAt         string                     `json:"indexedAt" cborgen:"indexedAt"`
	Name              string                     `json:"name" cborgen:"name"`
	Purpose           *string                    `json:"purpose" cborgen:"purpose"`
	Uri               string                     `json:"uri" cborgen:"uri"`
	Viewer            *GraphDefs_ListViewerState `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// GraphDefs_ListViewBasic is a "listViewBasic" in the app.bsky.graph.defs schema.
type GraphDefs_ListViewBasic struct {
	Avatar    *string                    `json:"avatar,omitempty" cborgen:"avatar,omitempty"`
	Cid       string                     `json:"cid" cborgen:"cid"`
	IndexedAt *string                    `json:"indexedAt,omitempty" cborgen:"indexedAt,omitempty"`
	Name      string                     `json:"name" cborgen:"name"`
	Purpose   *string                    `json:"purpose" cborgen:"purpose"`
	Uri       string                     `json:"uri" cborgen:"uri"`
	Viewer    *GraphDefs_ListViewerState `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// GraphDefs_ListViewerState is a "listViewerState" in the app.bsky.graph.defs schema.
type GraphDefs_ListViewerState struct {
	Muted *bool `json:"muted,omitempty" cborgen:"muted,omitempty"`
}
