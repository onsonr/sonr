import {RiArticleLine} from 'react-icons/ri';
import type {StructureBuilder} from 'sanity/desk';

export const PostMenuItem = (S: StructureBuilder) =>
	S.listItem()
		.title('Posts')
		.icon(RiArticleLine)
		.child(S.documentTypeList('post').title('Posts').filter('_type == $type').params({type: 'post'}));
