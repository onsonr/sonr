import {RiSettings5Line} from 'react-icons/ri';
import type {StructureBuilder} from 'sanity/desk';

export const SiteSettings = (S: StructureBuilder) =>
	S.listItem()
		.title('Global Settings')
		.icon(RiSettings5Line)
		.child(S.editor().schemaType('siteSettings').documentId('siteSettings'));
