const slTypes = 'sl-button, sl-checkbox, sl-color-picker, sl-input, sl-radio-group, sl-range, sl-rating, sl-select, sl-switch, sl-textarea'

/* Lightly modified version of the same function in htmx.js */
function shouldInclude(elt) {

	// sl-rating doesn't have a name attribute exposed through the Shoelace API
	if (elt.tagName == 'SL-RATING' && elt.getAttribute('name')) {
		return true;
	}
	if (elt.name === "" || elt.name == null || elt.disabled) {
		return false;
	}
	if (elt.type === "submit") {
		return false;
	}
	if (elt.tagName === 'SL-CHECKBOX' || elt.tagName === 'SL-SWITCH') {
		return elt.checked;
	}
	return true;
}

htmx.defineExtension('shoelace', {
	onEvent: function (name, evt) {
		if ((name === "htmx:configRequest") && (evt.detail.elt.tagName == 'FORM')) {
			evt.detail.elt.querySelectorAll(slTypes).forEach((elt) => {
				if (shouldInclude(elt)) {

					if (elt.tagName === 'SL-CHECKBOX' || elt.tagName === 'SL-SWITCH') {
						evt.detail.parameters[elt.name] = elt.checked

					} else if (elt.tagName == 'SL-RATING') {
						evt.detail.parameters[elt.getAttribute('name')] = elt.value

					} else {
						evt.detail.parameters[elt.name] = elt.value
					}
				}
			})
		}
	}
})
