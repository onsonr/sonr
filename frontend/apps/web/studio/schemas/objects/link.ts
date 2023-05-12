import {defineField} from 'sanity';

const link = defineField({
	title: 'URL',
	name: 'link',
	type: 'object',
	hidden: true,
	fields: [
		{
			title: 'URL',
			name: 'href',
			type: 'url',
			validation: (Rule) =>
				Rule.uri({
					allowRelative: true,
					scheme: ['https', 'http', 'mailto', 'tel']
				})
		}
	]
});

export default link;
