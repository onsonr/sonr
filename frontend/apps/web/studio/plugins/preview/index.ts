import type {DefaultDocumentNodeResolver} from 'sanity/desk';
import PostPreview from './post-preview';
import PagePreview from './page-preview';

export const previewDocumentNode = (): DefaultDocumentNodeResolver => {
	return (S, {schemaType}) => {
		if (schemaType === 'post') {
			return S.document().views([S.view.form(), S.view.component(PostPreview).title('Preview Post')]);
		}

		if (schemaType === 'page') {
			return S.document().views([S.view.form(), S.view.component(PagePreview).title('Preview Page')]);
		}

		return S.document().views([S.view.form()]);
	};
};
