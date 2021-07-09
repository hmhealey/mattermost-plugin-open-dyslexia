import * as React from 'react';
import {connect} from 'react-redux';

import {GlobalState} from 'mattermost-redux/types/store';

import {isFontEnabled} from '../../preferences';

const css = `
@font-face {
    font-family: "OpenDyslexic";
    src: url("/static/plugins/ca.hmhealey.plugin-open-dyslexic/fonts/OpenDyslexic-Regular.otf") format("opentype");
}

@font-face {
    font-family: "OpenDyslexicMono";
    src: url("/static/plugins/ca.hmhealey.plugin-open-dyslexic/fonts/OpenDyslexicMono-Regular.otf") format("opentype");
}

body,
.app__body.font--open_sans,
#SidebarContainer .SidebarChannelGroupHeader_groupButton {
    font-family: OpenDyslexic, "Open Sans", sans-serif;
}

code {
    font-family: OpenDyslexicMono, Menlo, Monaco, Consolas, "Courier New", monospace;
}
`;

type Props = {
    enabled: boolean;
}

function RootStyles(props: Props) {
    if (!props.enabled) {
        return null;
    }

    return (
        <style>{css}</style>
    );
}

function mapStateToProps(state: GlobalState) {
    return {
        enabled: isFontEnabled(state),
    };
}

export default connect(mapStateToProps)(RootStyles);
