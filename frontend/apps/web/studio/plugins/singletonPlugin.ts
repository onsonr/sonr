import {definePlugin} from 'sanity';

type SingletonPluginOptions = {
	types: string[];
};

export const singletonPlugin = definePlugin<SingletonPluginOptions>(({types}) => {
	return {
		name: 'singletonPlugin',
		document: {
			// Hide 'Singletons (such as Home)' from new document options
			// https://user-images.githubusercontent.com/81981/195728798-e0c6cf7e-d442-4e58-af3a-8cd99d7fcc28.png
			newDocumentOptions: (prev, {creationContext}) => {
				if (creationContext.type === 'global') {
					return prev.filter((templateItem) => !types.includes(templateItem.templateId));
				}

				return prev;
			},
			// Removes the "duplicate" action on the Singletons
			actions: (prev, {schemaType}) => {
				if (types.includes(schemaType)) {
					return prev.filter(
						({action}) => action !== 'duplicate' && action !== 'delete' && action !== 'unpublish'
					);
				}

				return prev;
			}
		}
	};
});
