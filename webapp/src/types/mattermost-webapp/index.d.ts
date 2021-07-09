export interface PluginRegistry {
    registerRootComponent(component: React.ElementType)

    // Add more if needed from https://developers.mattermost.com/extend/plugins/webapp/reference
}
