import $ from 'jquery';

export function filtering(trajectory) {
  let filtered_trajectory;

  $.ajax({
    url: 'http://localhost:8080/filter',
    method: 'post',
    async: false,
    data: JSON.stringify(trajectory),
  })
  .done(function(result) {
    filtered_trajectory = result;
  });

  return filtered_trajectory;
}
