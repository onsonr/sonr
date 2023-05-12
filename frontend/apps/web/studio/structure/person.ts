import {RiUserSmileLine} from 'react-icons/ri';
import type {StructureBuilder} from 'sanity/desk';

export const PersonMenuItem = (S: StructureBuilder) =>
	S.listItem()
		.title('Persons')
		.icon(RiUserSmileLine)
		.child(S.documentTypeList('person').title('Persons').filter('_type == $type').params({type: 'person'}));
