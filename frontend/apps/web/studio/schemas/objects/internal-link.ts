import {RiLinksLine} from 'react-icons/ri';
import {defineField} from 'sanity';

const internalLink = defineField({
	title: 'Internal Link',
	name: 'internalLink',
	type: 'object',
	hidden: true,
	icon: RiLinksLine,
	fields: [
		{
			name: 'title',
			title: 'Title',
			type: 'string',
			validation: (Rule) => Rule.required()
		},
		{
			name: 'link',
			title: 'Link',
			type: 'reference',
			to: [
				{
					type: 'page'
				}
			]
		}
	]
});

export default internalLink;;
