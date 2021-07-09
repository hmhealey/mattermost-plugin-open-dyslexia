import {Store, Action} from 'redux';

import {GlobalState} from 'mattermost-redux/types/store';

import RootStyles from './components/root_styles';

import manifest from './manifest';

import {PluginRegistry} from './types/mattermost-webapp';

export default class Plugin {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    public async initialize(registry: PluginRegistry, store: Store<GlobalState, Action<Record<string, unknown>>>) {
        registry.registerRootComponent(RootStyles);
    }
}

declare global {
    interface Window {
        registerPlugin(id: string, plugin: Plugin): void
    }
}

window.registerPlugin(manifest.id, new Plugin());
