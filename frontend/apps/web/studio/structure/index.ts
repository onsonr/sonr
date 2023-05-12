import type {StructureResolver} from 'sanity/desk';
import {SiteSettings} from './global-settings';
import {PageMenuItem} from './pages';
import {PersonMenuItem} from './person';
import {PostMenuItem} from './posts';
// import PostPreview from '../components/previews/post/post-preview';
// import PagePreview from '../components/previews/page/page-preview';

export const structure: StructureResolver = (S, _context) =>
	S.list()
		.title('Content')
		.items([SiteSettings(S), PageMenuItem(S), PostMenuItem(S), PersonMenuItem(S)]);

export const defaultDocumentNode = (S: any, props: any) => {
	const {schemaType} = props;

	if (schemaType === 'post') {
		// return S.document().views([S.view.form(), S.view.component(PostPreview).title('Preview Post')]);
	}

	if (schemaType === 'page') {
		// return S.document().views([S.view.form(), S.view.component(PagePreview).title('Preview Page')]);
	}

	return S.document().views([S.view.form()]);
};

export default structure;
