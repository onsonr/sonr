import {RiSettings5Line} from 'react-icons/ri';
import {defineType, defineField} from 'sanity';

const siteSettings = defineType({
	name: 'siteSettings',
	type: 'document',
	title: 'Site Settings',
	icon: RiSettings5Line,
	groups: [
		{
			name: 'meta',
			title: 'General Information'
		},
		{
			name: 'navigation',
			title: 'Main Navigation'
		},
		{
			name: 'social',
			title: 'Social'
		}
	],
	fields: [
		defineField({
			name: 'title',
			title: 'Title',
			type: 'string',
			description: 'Title of the page',
			group: 'meta',
			validation: (Rule) => Rule.required()
		}),
		defineField({
			name: 'description',
			title: 'Description',
			type: 'text',
			description: 'Description for search engines and social media.',
			group: 'meta',
			validation: (Rule) => Rule.required()
		}),
		defineField({
			type: 'metaFields',
			title: 'Meta',
			name: 'meta',
			group: 'meta'
		}),
		defineField({
			title: 'Navigation',
			name: 'navigation',
			description: 'Select pages or link for main navigation',
			group: 'navigation',
			type: 'array',
			of: [
				{
					title: 'Internal Link',
					type: 'internalLink'
				},
				{
					title: 'External Link',
					type: 'externalLink'
				}
			]
		}),
		defineField({
			name: 'socialFields',
			type: 'socialFields',
			description: 'Social media',
			group: 'social'
		})
	],
	preview: {
		select: {},
		prepare() {
			return {
				title: 'Global Settings'
			};
		}
	}
});

export default siteSettings;
