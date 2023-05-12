import {RiPagesLine} from 'react-icons/ri';
import type {StructureBuilder} from 'sanity/desk';

export const PageMenuItem = (S: StructureBuilder) =>
	S.listItem()
		.title('Pages')
		.icon(RiPagesLine)
		.child(S.documentTypeList('page').title('Pages').filter('_type == $type').params({type: 'page'}));
