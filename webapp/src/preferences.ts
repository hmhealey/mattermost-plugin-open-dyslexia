import {savePreferences} from 'mattermost-redux/actions/preferences';
import {getBool} from 'mattermost-redux/selectors/entities/preferences';
import {getCurrentUserId} from 'mattermost-redux/selectors/entities/users';
import {DispatchFunc, GetStateFunc} from 'mattermost-redux/types/actions';
import {GlobalState} from 'mattermost-redux/types/store';

const preferenceCategory = 'opendyslexic';
const preferenceName = 'enabled';

export function isFontEnabled(state: GlobalState) {
    return getBool(state, preferenceCategory, preferenceName, false);
}

export function setFontEnabled(enabled: boolean) {
    return (dispatch: DispatchFunc, getState: GetStateFunc) => {
        const currentUserId = getCurrentUserId(getState());

        dispatch(savePreferences(currentUserId, [{
            user_id: currentUserId,
            category: preferenceCategory,
            name: preferenceName,
            value: enabled.toString(),
        }]));
    };
}
