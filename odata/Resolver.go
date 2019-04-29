package odata

import "context"

type GqlResolver struct {
	Client *Client
}

func (r *GqlResolver) AccumulationRegisterPartiiTovarovVProizvodstve(ctx context.Context, args AccumulationRegisterPartiiTovarovVProizvodstveArgs) (*AccumulationRegisterPartiiTovarovVProizvodstve, error) {
	return r.Client.AccumulationRegisterPartiiTovarovVProizvodstve(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterPartiiTovarovVProizvodstves(ctx context.Context, args AccumulationRegisterPartiiTovarovVProizvodstvesArgs) (*[]AccumulationRegisterPartiiTovarovVProizvodstve, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterPartiiTovarovVProizvodstves(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterPartiiTovarovVProizvodstves(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterPartiiTovarovVProizvodstves(Where{})
}
func (r *GqlResolver) AccumulationRegisterPartiiTovarovVProizvodstveRecordType(ctx context.Context, args AccumulationRegisterPartiiTovarovVProizvodstveRecordTypeArgs) (*AccumulationRegisterPartiiTovarovVProizvodstveRecordType, error) {
	return r.Client.AccumulationRegisterPartiiTovarovVProizvodstveRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterPartiiTovarovVProizvodstveRecordTypes(ctx context.Context, args AccumulationRegisterPartiiTovarovVProizvodstveRecordTypesArgs) (*[]AccumulationRegisterPartiiTovarovVProizvodstveRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterPartiiTovarovVProizvodstveRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterPartiiTovarovVProizvodstveRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterPartiiTovarovVProizvodstveRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterVzaimoraschetySPodotchetnymiLitsami(ctx context.Context, args AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiArgs) (*AccumulationRegisterVzaimoraschetySPodotchetnymiLitsami, error) {
	return r.Client.AccumulationRegisterVzaimoraschetySPodotchetnymiLitsami(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamis(ctx context.Context, args AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamisArgs) (*[]AccumulationRegisterVzaimoraschetySPodotchetnymiLitsami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamis(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamis(Where{})
}
func (r *GqlResolver) AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType(ctx context.Context, args AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordTypeArgs) (*AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType, error) {
	return r.Client.AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordTypes(ctx context.Context, args AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordTypesArgs) (*[]AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterVzaimoraschetySPodotchetnymiLitsamiRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterVnutrennieZakazy(ctx context.Context, args AccumulationRegisterVnutrennieZakazyArgs) (*AccumulationRegisterVnutrennieZakazy, error) {
	return r.Client.AccumulationRegisterVnutrennieZakazy(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterVnutrennieZakazys(ctx context.Context, args AccumulationRegisterVnutrennieZakazysArgs) (*[]AccumulationRegisterVnutrennieZakazy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterVnutrennieZakazys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterVnutrennieZakazys(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterVnutrennieZakazys(Where{})
}
func (r *GqlResolver) AccumulationRegisterVnutrennieZakazyRecordType(ctx context.Context, args AccumulationRegisterVnutrennieZakazyRecordTypeArgs) (*AccumulationRegisterVnutrennieZakazyRecordType, error) {
	return r.Client.AccumulationRegisterVnutrennieZakazyRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterVnutrennieZakazyRecordTypes(ctx context.Context, args AccumulationRegisterVnutrennieZakazyRecordTypesArgs) (*[]AccumulationRegisterVnutrennieZakazyRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterVnutrennieZakazyRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterVnutrennieZakazyRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterVnutrennieZakazyRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKomitenta(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKomitentaArgs) (*AccumulationRegisterDenezhnyeSredstvaKomitenta, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKomitenta(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKomitentas(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKomitentasArgs) (*[]AccumulationRegisterDenezhnyeSredstvaKomitenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKomitentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKomitentas(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKomitentas(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKomitentaRecordType(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKomitentaRecordTypeArgs) (*AccumulationRegisterDenezhnyeSredstvaKomitentaRecordType, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKomitentaRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKomitentaRecordTypes(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKomitentaRecordTypesArgs) (*[]AccumulationRegisterDenezhnyeSredstvaKomitentaRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKomitentaRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKomitentaRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKomitentaRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterZakazyKlientov(ctx context.Context, args AccumulationRegisterZakazyKlientovArgs) (*AccumulationRegisterZakazyKlientov, error) {
	return r.Client.AccumulationRegisterZakazyKlientov(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterZakazyKlientovs(ctx context.Context, args AccumulationRegisterZakazyKlientovsArgs) (*[]AccumulationRegisterZakazyKlientov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterZakazyKlientovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterZakazyKlientovs(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterZakazyKlientovs(Where{})
}
func (r *GqlResolver) AccumulationRegisterZakazyKlientovRecordType(ctx context.Context, args AccumulationRegisterZakazyKlientovRecordTypeArgs) (*AccumulationRegisterZakazyKlientovRecordType, error) {
	return r.Client.AccumulationRegisterZakazyKlientovRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterZakazyKlientovRecordTypes(ctx context.Context, args AccumulationRegisterZakazyKlientovRecordTypesArgs) (*[]AccumulationRegisterZakazyKlientovRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterZakazyKlientovRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterZakazyKlientovRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterZakazyKlientovRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyPoFinmonitoringuRoznitsa(ctx context.Context, args AccumulationRegisterSummyPoFinmonitoringuRoznitsaArgs) (*AccumulationRegisterSummyPoFinmonitoringuRoznitsa, error) {
	return r.Client.AccumulationRegisterSummyPoFinmonitoringuRoznitsa(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyPoFinmonitoringuRoznitsas(ctx context.Context, args AccumulationRegisterSummyPoFinmonitoringuRoznitsasArgs) (*[]AccumulationRegisterSummyPoFinmonitoringuRoznitsa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyPoFinmonitoringuRoznitsas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyPoFinmonitoringuRoznitsas(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyPoFinmonitoringuRoznitsas(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType(ctx context.Context, args AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordTypeArgs) (*AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType, error) {
	return r.Client.AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordTypes(ctx context.Context, args AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordTypesArgs) (*[]AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyPoFinmonitoringuRoznitsaRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKPolucheniiu(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKPolucheniiuArgs) (*AccumulationRegisterDenezhnyeSredstvaKPolucheniiu, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKPolucheniiu(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKPolucheniius(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKPolucheniiusArgs) (*[]AccumulationRegisterDenezhnyeSredstvaKPolucheniiu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKPolucheniius(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKPolucheniius(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKPolucheniius(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordTypeArgs) (*AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordTypes(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordTypesArgs) (*[]AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKPolucheniiuRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterProdazhiPoDiskontnymKartam(ctx context.Context, args AccumulationRegisterProdazhiPoDiskontnymKartamArgs) (*AccumulationRegisterProdazhiPoDiskontnymKartam, error) {
	return r.Client.AccumulationRegisterProdazhiPoDiskontnymKartam(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterProdazhiPoDiskontnymKartams(ctx context.Context, args AccumulationRegisterProdazhiPoDiskontnymKartamsArgs) (*[]AccumulationRegisterProdazhiPoDiskontnymKartam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterProdazhiPoDiskontnymKartams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterProdazhiPoDiskontnymKartams(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterProdazhiPoDiskontnymKartams(Where{})
}
func (r *GqlResolver) AccumulationRegisterProdazhiPoDiskontnymKartamRecordType(ctx context.Context, args AccumulationRegisterProdazhiPoDiskontnymKartamRecordTypeArgs) (*AccumulationRegisterProdazhiPoDiskontnymKartamRecordType, error) {
	return r.Client.AccumulationRegisterProdazhiPoDiskontnymKartamRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterProdazhiPoDiskontnymKartamRecordTypes(ctx context.Context, args AccumulationRegisterProdazhiPoDiskontnymKartamRecordTypesArgs) (*[]AccumulationRegisterProdazhiPoDiskontnymKartamRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterProdazhiPoDiskontnymKartamRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterProdazhiPoDiskontnymKartamRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterProdazhiPoDiskontnymKartamRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryPoluchennye(ctx context.Context, args AccumulationRegisterTovaryPoluchennyeArgs) (*AccumulationRegisterTovaryPoluchennye, error) {
	return r.Client.AccumulationRegisterTovaryPoluchennye(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryPoluchennyes(ctx context.Context, args AccumulationRegisterTovaryPoluchennyesArgs) (*[]AccumulationRegisterTovaryPoluchennye, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryPoluchennyes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryPoluchennyes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryPoluchennyes(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryPoluchennyeRecordType(ctx context.Context, args AccumulationRegisterTovaryPoluchennyeRecordTypeArgs) (*AccumulationRegisterTovaryPoluchennyeRecordType, error) {
	return r.Client.AccumulationRegisterTovaryPoluchennyeRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryPoluchennyeRecordTypes(ctx context.Context, args AccumulationRegisterTovaryPoluchennyeRecordTypesArgs) (*[]AccumulationRegisterTovaryPoluchennyeRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryPoluchennyeRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryPoluchennyeRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryPoluchennyeRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterSvobodnyeOstatki(ctx context.Context, args AccumulationRegisterSvobodnyeOstatkiArgs) (*AccumulationRegisterSvobodnyeOstatki, error) {
	return r.Client.AccumulationRegisterSvobodnyeOstatki(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSvobodnyeOstatkis(ctx context.Context, args AccumulationRegisterSvobodnyeOstatkisArgs) (*[]AccumulationRegisterSvobodnyeOstatki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSvobodnyeOstatkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSvobodnyeOstatkis(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSvobodnyeOstatkis(Where{})
}
func (r *GqlResolver) AccumulationRegisterSvobodnyeOstatkiRecordType(ctx context.Context, args AccumulationRegisterSvobodnyeOstatkiRecordTypeArgs) (*AccumulationRegisterSvobodnyeOstatkiRecordType, error) {
	return r.Client.AccumulationRegisterSvobodnyeOstatkiRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSvobodnyeOstatkiRecordTypes(ctx context.Context, args AccumulationRegisterSvobodnyeOstatkiRecordTypesArgs) (*[]AccumulationRegisterSvobodnyeOstatkiRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSvobodnyeOstatkiRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSvobodnyeOstatkiRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSvobodnyeOstatkiRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyVRassrochku(ctx context.Context, args AccumulationRegisterSummyVRassrochkuArgs) (*AccumulationRegisterSummyVRassrochku, error) {
	return r.Client.AccumulationRegisterSummyVRassrochku(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyVRassrochkus(ctx context.Context, args AccumulationRegisterSummyVRassrochkusArgs) (*[]AccumulationRegisterSummyVRassrochku, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyVRassrochkus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyVRassrochkus(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyVRassrochkus(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyVRassrochkuRecordType(ctx context.Context, args AccumulationRegisterSummyVRassrochkuRecordTypeArgs) (*AccumulationRegisterSummyVRassrochkuRecordType, error) {
	return r.Client.AccumulationRegisterSummyVRassrochkuRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyVRassrochkuRecordTypes(ctx context.Context, args AccumulationRegisterSummyVRassrochkuRecordTypesArgs) (*[]AccumulationRegisterSummyVRassrochkuRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyVRassrochkuRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyVRassrochkuRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyVRassrochkuRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterGrafikPlatezhei(ctx context.Context, args AccumulationRegisterGrafikPlatezheiArgs) (*AccumulationRegisterGrafikPlatezhei, error) {
	return r.Client.AccumulationRegisterGrafikPlatezhei(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterGrafikPlatezheis(ctx context.Context, args AccumulationRegisterGrafikPlatezheisArgs) (*[]AccumulationRegisterGrafikPlatezhei, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterGrafikPlatezheis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterGrafikPlatezheis(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterGrafikPlatezheis(Where{})
}
func (r *GqlResolver) AccumulationRegisterGrafikPlatezheiRecordType(ctx context.Context, args AccumulationRegisterGrafikPlatezheiRecordTypeArgs) (*AccumulationRegisterGrafikPlatezheiRecordType, error) {
	return r.Client.AccumulationRegisterGrafikPlatezheiRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterGrafikPlatezheiRecordTypes(ctx context.Context, args AccumulationRegisterGrafikPlatezheiRecordTypesArgs) (*[]AccumulationRegisterGrafikPlatezheiRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterGrafikPlatezheiRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterGrafikPlatezheiRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterGrafikPlatezheiRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterRoznichnaiaVyruchka(ctx context.Context, args AccumulationRegisterRoznichnaiaVyruchkaArgs) (*AccumulationRegisterRoznichnaiaVyruchka, error) {
	return r.Client.AccumulationRegisterRoznichnaiaVyruchka(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterRoznichnaiaVyruchkas(ctx context.Context, args AccumulationRegisterRoznichnaiaVyruchkasArgs) (*[]AccumulationRegisterRoznichnaiaVyruchka, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterRoznichnaiaVyruchkas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterRoznichnaiaVyruchkas(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterRoznichnaiaVyruchkas(Where{})
}
func (r *GqlResolver) AccumulationRegisterRoznichnaiaVyruchkaRecordType(ctx context.Context, args AccumulationRegisterRoznichnaiaVyruchkaRecordTypeArgs) (*AccumulationRegisterRoznichnaiaVyruchkaRecordType, error) {
	return r.Client.AccumulationRegisterRoznichnaiaVyruchkaRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterRoznichnaiaVyruchkaRecordTypes(ctx context.Context, args AccumulationRegisterRoznichnaiaVyruchkaRecordTypesArgs) (*[]AccumulationRegisterRoznichnaiaVyruchkaRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterRoznichnaiaVyruchkaRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterRoznichnaiaVyruchkaRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterRoznichnaiaVyruchkaRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryVPuti(ctx context.Context, args AccumulationRegisterTovaryVPutiArgs) (*AccumulationRegisterTovaryVPuti, error) {
	return r.Client.AccumulationRegisterTovaryVPuti(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryVPutis(ctx context.Context, args AccumulationRegisterTovaryVPutisArgs) (*[]AccumulationRegisterTovaryVPuti, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryVPutis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryVPutis(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryVPutis(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryVPutiRecordType(ctx context.Context, args AccumulationRegisterTovaryVPutiRecordTypeArgs) (*AccumulationRegisterTovaryVPutiRecordType, error) {
	return r.Client.AccumulationRegisterTovaryVPutiRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryVPutiRecordTypes(ctx context.Context, args AccumulationRegisterTovaryVPutiRecordTypesArgs) (*[]AccumulationRegisterTovaryVPutiRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryVPutiRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryVPutiRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryVPutiRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterPoteriMetallaVProizvodstve(ctx context.Context, args AccumulationRegisterPoteriMetallaVProizvodstveArgs) (*AccumulationRegisterPoteriMetallaVProizvodstve, error) {
	return r.Client.AccumulationRegisterPoteriMetallaVProizvodstve(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterPoteriMetallaVProizvodstves(ctx context.Context, args AccumulationRegisterPoteriMetallaVProizvodstvesArgs) (*[]AccumulationRegisterPoteriMetallaVProizvodstve, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterPoteriMetallaVProizvodstves(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterPoteriMetallaVProizvodstves(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterPoteriMetallaVProizvodstves(Where{})
}
func (r *GqlResolver) AccumulationRegisterPoteriMetallaVProizvodstveRecordType(ctx context.Context, args AccumulationRegisterPoteriMetallaVProizvodstveRecordTypeArgs) (*AccumulationRegisterPoteriMetallaVProizvodstveRecordType, error) {
	return r.Client.AccumulationRegisterPoteriMetallaVProizvodstveRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterPoteriMetallaVProizvodstveRecordTypes(ctx context.Context, args AccumulationRegisterPoteriMetallaVProizvodstveRecordTypesArgs) (*[]AccumulationRegisterPoteriMetallaVProizvodstveRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterPoteriMetallaVProizvodstveRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterPoteriMetallaVProizvodstveRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterPoteriMetallaVProizvodstveRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterPartiiTovarovNaSkladakh(ctx context.Context, args AccumulationRegisterPartiiTovarovNaSkladakhArgs) (*AccumulationRegisterPartiiTovarovNaSkladakh, error) {
	return r.Client.AccumulationRegisterPartiiTovarovNaSkladakh(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterPartiiTovarovNaSkladakhs(ctx context.Context, args AccumulationRegisterPartiiTovarovNaSkladakhsArgs) (*[]AccumulationRegisterPartiiTovarovNaSkladakh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterPartiiTovarovNaSkladakhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterPartiiTovarovNaSkladakhs(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterPartiiTovarovNaSkladakhs(Where{})
}
func (r *GqlResolver) ProductActionDocument(ctx context.Context, args ProductActionDocumentArgs) (*ProductActionDocument, error) {
	return r.Client.ProductActionDocument(args.Key, nil)
}
func (r *GqlResolver) ProductActionDocuments(ctx context.Context, args ProductActionDocumentsArgs) (*[]ProductActionDocument, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.ProductActionDocuments(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.ProductActionDocuments(Where{Filter: *args.Filter})
	}
	return r.Client.ProductActionDocuments(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyDokumentovDliaObmena(ctx context.Context, args AccumulationRegisterSummyDokumentovDliaObmenaArgs) (*AccumulationRegisterSummyDokumentovDliaObmena, error) {
	return r.Client.AccumulationRegisterSummyDokumentovDliaObmena(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyDokumentovDliaObmenas(ctx context.Context, args AccumulationRegisterSummyDokumentovDliaObmenasArgs) (*[]AccumulationRegisterSummyDokumentovDliaObmena, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyDokumentovDliaObmenas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyDokumentovDliaObmenas(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyDokumentovDliaObmenas(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyDokumentovDliaObmenaRecordType(ctx context.Context, args AccumulationRegisterSummyDokumentovDliaObmenaRecordTypeArgs) (*AccumulationRegisterSummyDokumentovDliaObmenaRecordType, error) {
	return r.Client.AccumulationRegisterSummyDokumentovDliaObmenaRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyDokumentovDliaObmenaRecordTypes(ctx context.Context, args AccumulationRegisterSummyDokumentovDliaObmenaRecordTypesArgs) (*[]AccumulationRegisterSummyDokumentovDliaObmenaRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyDokumentovDliaObmenaRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyDokumentovDliaObmenaRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyDokumentovDliaObmenaRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterDvizheniiaDenezhnykhSredstv(ctx context.Context, args AccumulationRegisterDvizheniiaDenezhnykhSredstvArgs) (*AccumulationRegisterDvizheniiaDenezhnykhSredstv, error) {
	return r.Client.AccumulationRegisterDvizheniiaDenezhnykhSredstv(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDvizheniiaDenezhnykhSredstvs(ctx context.Context, args AccumulationRegisterDvizheniiaDenezhnykhSredstvsArgs) (*[]AccumulationRegisterDvizheniiaDenezhnykhSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDvizheniiaDenezhnykhSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDvizheniiaDenezhnykhSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDvizheniiaDenezhnykhSredstvs(Where{})
}
func (r *GqlResolver) AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType(ctx context.Context, args AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordTypeArgs) (*AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType, error) {
	return r.Client.AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordTypes(ctx context.Context, args AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordTypesArgs) (*[]AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDvizheniiaDenezhnykhSredstvRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterProdazhiPoStatiam(ctx context.Context, args AccumulationRegisterProdazhiPoStatiamArgs) (*AccumulationRegisterProdazhiPoStatiam, error) {
	return r.Client.AccumulationRegisterProdazhiPoStatiam(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterProdazhiPoStatiams(ctx context.Context, args AccumulationRegisterProdazhiPoStatiamsArgs) (*[]AccumulationRegisterProdazhiPoStatiam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterProdazhiPoStatiams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterProdazhiPoStatiams(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterProdazhiPoStatiams(Where{})
}
func (r *GqlResolver) AccumulationRegisterProdazhiPoStatiamRecordType(ctx context.Context, args AccumulationRegisterProdazhiPoStatiamRecordTypeArgs) (*AccumulationRegisterProdazhiPoStatiamRecordType, error) {
	return r.Client.AccumulationRegisterProdazhiPoStatiamRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterProdazhiPoStatiamRecordTypes(ctx context.Context, args AccumulationRegisterProdazhiPoStatiamRecordTypesArgs) (*[]AccumulationRegisterProdazhiPoStatiamRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterProdazhiPoStatiamRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterProdazhiPoStatiamRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterProdazhiPoStatiamRecordTypes(Where{})
}
func (r *GqlResolver) InformationRegisterTsenyNomenklatury(ctx context.Context, args InformationRegisterTsenyNomenklaturyArgs) (*InformationRegisterTsenyNomenklatury, error) {
	return r.Client.InformationRegisterTsenyNomenklatury(args.Key, nil)
}
func (r *GqlResolver) InformationRegisterTsenyNomenklaturys(ctx context.Context, args InformationRegisterTsenyNomenklaturysArgs) (*[]InformationRegisterTsenyNomenklatury, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.InformationRegisterTsenyNomenklaturys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.InformationRegisterTsenyNomenklaturys(Where{Filter: *args.Filter})
	}
	return r.Client.InformationRegisterTsenyNomenklaturys(Where{})
}
func (r *GqlResolver) InformationRegisterTsenyNomenklaturyRecordType(ctx context.Context, args InformationRegisterTsenyNomenklaturyRecordTypeArgs) (*InformationRegisterTsenyNomenklaturyRecordType, error) {
	return r.Client.InformationRegisterTsenyNomenklaturyRecordType(args.Key, nil)
}
func (r *GqlResolver) InformationRegisterTsenyNomenklaturyRecordTypes(ctx context.Context, args InformationRegisterTsenyNomenklaturyRecordTypesArgs) (*[]InformationRegisterTsenyNomenklaturyRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.InformationRegisterTsenyNomenklaturyRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.InformationRegisterTsenyNomenklaturyRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.InformationRegisterTsenyNomenklaturyRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse(ctx context.Context, args AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseArgs) (*AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse, error) {
	return r.Client.AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitses(ctx context.Context, args AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitsesArgs) (*[]AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitse, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitses(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitses(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitses(Where{})
}
func (r *GqlResolver) AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType(ctx context.Context, args AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordTypeArgs) (*AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType, error) {
	return r.Client.AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordTypes(ctx context.Context, args AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordTypesArgs) (*[]AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSvodnyeDannyePoProdazhamVRoznitseRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaVRezerve(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaVRezerveArgs) (*AccumulationRegisterDenezhnyeSredstvaVRezerve, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaVRezerve(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaVRezerves(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaVRezervesArgs) (*[]AccumulationRegisterDenezhnyeSredstvaVRezerve, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaVRezerves(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaVRezerves(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaVRezerves(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaVRezerveRecordType(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaVRezerveRecordTypeArgs) (*AccumulationRegisterDenezhnyeSredstvaVRezerveRecordType, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaVRezerveRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaVRezerveRecordTypes(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaVRezerveRecordTypesArgs) (*[]AccumulationRegisterDenezhnyeSredstvaVRezerveRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaVRezerveRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaVRezerveRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaVRezerveRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh(ctx context.Context, args AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhArgs) (*AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh, error) {
	return r.Client.AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhs(ctx context.Context, args AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhsArgs) (*[]AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhs(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhs(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType(ctx context.Context, args AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordTypeArgs) (*AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType, error) {
	return r.Client.AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordTypes(ctx context.Context, args AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordTypesArgs) (*[]AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryVNeavtomatizirovannykhTorgovykhTochkakhRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterDavalcheskiiMetallPoteri(ctx context.Context, args AccumulationRegisterDavalcheskiiMetallPoteriArgs) (*AccumulationRegisterDavalcheskiiMetallPoteri, error) {
	return r.Client.AccumulationRegisterDavalcheskiiMetallPoteri(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDavalcheskiiMetallPoteris(ctx context.Context, args AccumulationRegisterDavalcheskiiMetallPoterisArgs) (*[]AccumulationRegisterDavalcheskiiMetallPoteri, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDavalcheskiiMetallPoteris(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDavalcheskiiMetallPoteris(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDavalcheskiiMetallPoteris(Where{})
}
func (r *GqlResolver) AccumulationRegisterDavalcheskiiMetallPoteriRecordType(ctx context.Context, args AccumulationRegisterDavalcheskiiMetallPoteriRecordTypeArgs) (*AccumulationRegisterDavalcheskiiMetallPoteriRecordType, error) {
	return r.Client.AccumulationRegisterDavalcheskiiMetallPoteriRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDavalcheskiiMetallPoteriRecordTypes(ctx context.Context, args AccumulationRegisterDavalcheskiiMetallPoteriRecordTypesArgs) (*[]AccumulationRegisterDavalcheskiiMetallPoteriRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDavalcheskiiMetallPoteriRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDavalcheskiiMetallPoteriRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDavalcheskiiMetallPoteriRecordTypes(Where{})
}
func (r *GqlResolver) InformationRegisterTsenyPoPreiskurantu(ctx context.Context, args InformationRegisterTsenyPoPreiskurantuArgs) (*InformationRegisterTsenyPoPreiskurantu, error) {
	return r.Client.InformationRegisterTsenyPoPreiskurantu(args.Key, nil)
}
func (r *GqlResolver) InformationRegisterTsenyPoPreiskurantus(ctx context.Context, args InformationRegisterTsenyPoPreiskurantusArgs) (*[]InformationRegisterTsenyPoPreiskurantu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.InformationRegisterTsenyPoPreiskurantus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.InformationRegisterTsenyPoPreiskurantus(Where{Filter: *args.Filter})
	}
	return r.Client.InformationRegisterTsenyPoPreiskurantus(Where{})
}
func (r *GqlResolver) InformationRegisterTsenyPoPreiskurantuRecordType(ctx context.Context, args InformationRegisterTsenyPoPreiskurantuRecordTypeArgs) (*InformationRegisterTsenyPoPreiskurantuRecordType, error) {
	return r.Client.InformationRegisterTsenyPoPreiskurantuRecordType(args.Key, nil)
}
func (r *GqlResolver) InformationRegisterTsenyPoPreiskurantuRecordTypes(ctx context.Context, args InformationRegisterTsenyPoPreiskurantuRecordTypesArgs) (*[]InformationRegisterTsenyPoPreiskurantuRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.InformationRegisterTsenyPoPreiskurantuRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.InformationRegisterTsenyPoPreiskurantuRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.InformationRegisterTsenyPoPreiskurantuRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryVOtbore(ctx context.Context, args AccumulationRegisterTovaryVOtboreArgs) (*AccumulationRegisterTovaryVOtbore, error) {
	return r.Client.AccumulationRegisterTovaryVOtbore(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryVOtbores(ctx context.Context, args AccumulationRegisterTovaryVOtboresArgs) (*[]AccumulationRegisterTovaryVOtbore, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryVOtbores(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryVOtbores(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryVOtbores(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryVOtboreRecordType(ctx context.Context, args AccumulationRegisterTovaryVOtboreRecordTypeArgs) (*AccumulationRegisterTovaryVOtboreRecordType, error) {
	return r.Client.AccumulationRegisterTovaryVOtboreRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryVOtboreRecordTypes(ctx context.Context, args AccumulationRegisterTovaryVOtboreRecordTypesArgs) (*[]AccumulationRegisterTovaryVOtboreRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryVOtboreRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryVOtboreRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryVOtboreRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterRealizovannyeTovary(ctx context.Context, args AccumulationRegisterRealizovannyeTovaryArgs) (*AccumulationRegisterRealizovannyeTovary, error) {
	return r.Client.AccumulationRegisterRealizovannyeTovary(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterRealizovannyeTovarys(ctx context.Context, args AccumulationRegisterRealizovannyeTovarysArgs) (*[]AccumulationRegisterRealizovannyeTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterRealizovannyeTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterRealizovannyeTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterRealizovannyeTovarys(Where{})
}
func (r *GqlResolver) AccumulationRegisterRealizovannyeTovaryRecordType(ctx context.Context, args AccumulationRegisterRealizovannyeTovaryRecordTypeArgs) (*AccumulationRegisterRealizovannyeTovaryRecordType, error) {
	return r.Client.AccumulationRegisterRealizovannyeTovaryRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterRealizovannyeTovaryRecordTypes(ctx context.Context, args AccumulationRegisterRealizovannyeTovaryRecordTypesArgs) (*[]AccumulationRegisterRealizovannyeTovaryRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterRealizovannyeTovaryRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterRealizovannyeTovaryRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterRealizovannyeTovaryRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKomissionera(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKomissioneraArgs) (*AccumulationRegisterDenezhnyeSredstvaKomissionera, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKomissionera(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKomissioneras(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKomissionerasArgs) (*[]AccumulationRegisterDenezhnyeSredstvaKomissionera, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKomissioneras(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKomissioneras(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKomissioneras(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordTypeArgs) (*AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordTypes(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordTypesArgs) (*[]AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKomissioneraRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterProdazhi1(ctx context.Context, args AccumulationRegisterProdazhi1Args) (*AccumulationRegisterProdazhi1, error) {
	return r.Client.AccumulationRegisterProdazhi1(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterProdazhi1s(ctx context.Context, args AccumulationRegisterProdazhi1sArgs) (*[]AccumulationRegisterProdazhi1, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterProdazhi1s(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterProdazhi1s(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterProdazhi1s(Where{})
}
func (r *GqlResolver) AccumulationRegisterProdazhi1RecordType(ctx context.Context, args AccumulationRegisterProdazhi1RecordTypeArgs) (*AccumulationRegisterProdazhi1RecordType, error) {
	return r.Client.AccumulationRegisterProdazhi1RecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterProdazhi1RecordTypes(ctx context.Context, args AccumulationRegisterProdazhi1RecordTypesArgs) (*[]AccumulationRegisterProdazhi1RecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterProdazhi1RecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterProdazhi1RecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterProdazhi1RecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryNaSkladakhAM(ctx context.Context, args AccumulationRegisterTovaryNaSkladakhAMArgs) (*AccumulationRegisterTovaryNaSkladakhAM, error) {
	return r.Client.AccumulationRegisterTovaryNaSkladakhAM(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryNaSkladakhAMs(ctx context.Context, args AccumulationRegisterTovaryNaSkladakhAMsArgs) (*[]AccumulationRegisterTovaryNaSkladakhAM, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryNaSkladakhAMs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryNaSkladakhAMs(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryNaSkladakhAMs(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryNaSkladakhAMRecordType(ctx context.Context, args AccumulationRegisterTovaryNaSkladakhAMRecordTypeArgs) (*AccumulationRegisterTovaryNaSkladakhAMRecordType, error) {
	return r.Client.AccumulationRegisterTovaryNaSkladakhAMRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryNaSkladakhAMRecordTypes(ctx context.Context, args AccumulationRegisterTovaryNaSkladakhAMRecordTypesArgs) (*[]AccumulationRegisterTovaryNaSkladakhAMRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryNaSkladakhAMRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryNaSkladakhAMRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryNaSkladakhAMRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyPoFinmonitoringu(ctx context.Context, args AccumulationRegisterSummyPoFinmonitoringuArgs) (*AccumulationRegisterSummyPoFinmonitoringu, error) {
	return r.Client.AccumulationRegisterSummyPoFinmonitoringu(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyPoFinmonitoringus(ctx context.Context, args AccumulationRegisterSummyPoFinmonitoringusArgs) (*[]AccumulationRegisterSummyPoFinmonitoringu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyPoFinmonitoringus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyPoFinmonitoringus(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyPoFinmonitoringus(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyPoFinmonitoringuRecordType(ctx context.Context, args AccumulationRegisterSummyPoFinmonitoringuRecordTypeArgs) (*AccumulationRegisterSummyPoFinmonitoringuRecordType, error) {
	return r.Client.AccumulationRegisterSummyPoFinmonitoringuRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyPoFinmonitoringuRecordTypes(ctx context.Context, args AccumulationRegisterSummyPoFinmonitoringuRecordTypesArgs) (*[]AccumulationRegisterSummyPoFinmonitoringuRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyPoFinmonitoringuRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyPoFinmonitoringuRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyPoFinmonitoringuRecordTypes(Where{})
}
func (r *GqlResolver) InformationRegisterTsenyNomenklaturyKontragentov(ctx context.Context, args InformationRegisterTsenyNomenklaturyKontragentovArgs) (*InformationRegisterTsenyNomenklaturyKontragentov, error) {
	return r.Client.InformationRegisterTsenyNomenklaturyKontragentov(args.Key, nil)
}
func (r *GqlResolver) InformationRegisterTsenyNomenklaturyKontragentovs(ctx context.Context, args InformationRegisterTsenyNomenklaturyKontragentovsArgs) (*[]InformationRegisterTsenyNomenklaturyKontragentov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.InformationRegisterTsenyNomenklaturyKontragentovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.InformationRegisterTsenyNomenklaturyKontragentovs(Where{Filter: *args.Filter})
	}
	return r.Client.InformationRegisterTsenyNomenklaturyKontragentovs(Where{})
}
func (r *GqlResolver) InformationRegisterTsenyNomenklaturyKontragentovRecordType(ctx context.Context, args InformationRegisterTsenyNomenklaturyKontragentovRecordTypeArgs) (*InformationRegisterTsenyNomenklaturyKontragentovRecordType, error) {
	return r.Client.InformationRegisterTsenyNomenklaturyKontragentovRecordType(args.Key, nil)
}
func (r *GqlResolver) InformationRegisterTsenyNomenklaturyKontragentovRecordTypes(ctx context.Context, args InformationRegisterTsenyNomenklaturyKontragentovRecordTypesArgs) (*[]InformationRegisterTsenyNomenklaturyKontragentovRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.InformationRegisterTsenyNomenklaturyKontragentovRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.InformationRegisterTsenyNomenklaturyKontragentovRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.InformationRegisterTsenyNomenklaturyKontragentovRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterVzaimoraschetySKontragentami(ctx context.Context, args AccumulationRegisterVzaimoraschetySKontragentamiArgs) (*AccumulationRegisterVzaimoraschetySKontragentami, error) {
	return r.Client.AccumulationRegisterVzaimoraschetySKontragentami(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterVzaimoraschetySKontragentamis(ctx context.Context, args AccumulationRegisterVzaimoraschetySKontragentamisArgs) (*[]AccumulationRegisterVzaimoraschetySKontragentami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterVzaimoraschetySKontragentamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterVzaimoraschetySKontragentamis(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterVzaimoraschetySKontragentamis(Where{})
}
func (r *GqlResolver) AccumulationRegisterVzaimoraschetySKontragentamiRecordType(ctx context.Context, args AccumulationRegisterVzaimoraschetySKontragentamiRecordTypeArgs) (*AccumulationRegisterVzaimoraschetySKontragentamiRecordType, error) {
	return r.Client.AccumulationRegisterVzaimoraschetySKontragentamiRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterVzaimoraschetySKontragentamiRecordTypes(ctx context.Context, args AccumulationRegisterVzaimoraschetySKontragentamiRecordTypesArgs) (*[]AccumulationRegisterVzaimoraschetySKontragentamiRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterVzaimoraschetySKontragentamiRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterVzaimoraschetySKontragentamiRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterVzaimoraschetySKontragentamiRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyPokupokPoDiskontnymKartam(ctx context.Context, args AccumulationRegisterSummyPokupokPoDiskontnymKartamArgs) (*AccumulationRegisterSummyPokupokPoDiskontnymKartam, error) {
	return r.Client.AccumulationRegisterSummyPokupokPoDiskontnymKartam(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyPokupokPoDiskontnymKartams(ctx context.Context, args AccumulationRegisterSummyPokupokPoDiskontnymKartamsArgs) (*[]AccumulationRegisterSummyPokupokPoDiskontnymKartam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyPokupokPoDiskontnymKartams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyPokupokPoDiskontnymKartams(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyPokupokPoDiskontnymKartams(Where{})
}
func (r *GqlResolver) AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType(ctx context.Context, args AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordTypeArgs) (*AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType, error) {
	return r.Client.AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordTypes(ctx context.Context, args AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordTypesArgs) (*[]AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterSummyPokupokPoDiskontnymKartamRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterVypolneniePlanaProdazh(ctx context.Context, args AccumulationRegisterVypolneniePlanaProdazhArgs) (*AccumulationRegisterVypolneniePlanaProdazh, error) {
	return r.Client.AccumulationRegisterVypolneniePlanaProdazh(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterVypolneniePlanaProdazhs(ctx context.Context, args AccumulationRegisterVypolneniePlanaProdazhsArgs) (*[]AccumulationRegisterVypolneniePlanaProdazh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterVypolneniePlanaProdazhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterVypolneniePlanaProdazhs(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterVypolneniePlanaProdazhs(Where{})
}
func (r *GqlResolver) AccumulationRegisterVypolneniePlanaProdazhRecordType(ctx context.Context, args AccumulationRegisterVypolneniePlanaProdazhRecordTypeArgs) (*AccumulationRegisterVypolneniePlanaProdazhRecordType, error) {
	return r.Client.AccumulationRegisterVypolneniePlanaProdazhRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterVypolneniePlanaProdazhRecordTypes(ctx context.Context, args AccumulationRegisterVypolneniePlanaProdazhRecordTypesArgs) (*[]AccumulationRegisterVypolneniePlanaProdazhRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterVypolneniePlanaProdazhRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterVypolneniePlanaProdazhRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterVypolneniePlanaProdazhRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterDavalcheskiiMetall(ctx context.Context, args AccumulationRegisterDavalcheskiiMetallArgs) (*AccumulationRegisterDavalcheskiiMetall, error) {
	return r.Client.AccumulationRegisterDavalcheskiiMetall(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDavalcheskiiMetalls(ctx context.Context, args AccumulationRegisterDavalcheskiiMetallsArgs) (*[]AccumulationRegisterDavalcheskiiMetall, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDavalcheskiiMetalls(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDavalcheskiiMetalls(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDavalcheskiiMetalls(Where{})
}
func (r *GqlResolver) AccumulationRegisterDavalcheskiiMetallRecordType(ctx context.Context, args AccumulationRegisterDavalcheskiiMetallRecordTypeArgs) (*AccumulationRegisterDavalcheskiiMetallRecordType, error) {
	return r.Client.AccumulationRegisterDavalcheskiiMetallRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDavalcheskiiMetallRecordTypes(ctx context.Context, args AccumulationRegisterDavalcheskiiMetallRecordTypesArgs) (*[]AccumulationRegisterDavalcheskiiMetallRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDavalcheskiiMetallRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDavalcheskiiMetallRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDavalcheskiiMetallRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstva(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaArgs) (*AccumulationRegisterDenezhnyeSredstva, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstva(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvas(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvasArgs) (*[]AccumulationRegisterDenezhnyeSredstva, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvas(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvas(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaRecordType(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaRecordTypeArgs) (*AccumulationRegisterDenezhnyeSredstvaRecordType, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaRecordTypes(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaRecordTypesArgs) (*[]AccumulationRegisterDenezhnyeSredstvaRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryPeredannye(ctx context.Context, args AccumulationRegisterTovaryPeredannyeArgs) (*AccumulationRegisterTovaryPeredannye, error) {
	return r.Client.AccumulationRegisterTovaryPeredannye(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryPeredannyes(ctx context.Context, args AccumulationRegisterTovaryPeredannyesArgs) (*[]AccumulationRegisterTovaryPeredannye, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryPeredannyes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryPeredannyes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryPeredannyes(Where{})
}
func (r *GqlResolver) AccumulationRegisterTovaryPeredannyeRecordType(ctx context.Context, args AccumulationRegisterTovaryPeredannyeRecordTypeArgs) (*AccumulationRegisterTovaryPeredannyeRecordType, error) {
	return r.Client.AccumulationRegisterTovaryPeredannyeRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterTovaryPeredannyeRecordTypes(ctx context.Context, args AccumulationRegisterTovaryPeredannyeRecordTypesArgs) (*[]AccumulationRegisterTovaryPeredannyeRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterTovaryPeredannyeRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterTovaryPeredannyeRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterTovaryPeredannyeRecordTypes(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKSpisaniiu(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKSpisaniiuArgs) (*AccumulationRegisterDenezhnyeSredstvaKSpisaniiu, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKSpisaniiu(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKSpisaniius(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKSpisaniiusArgs) (*[]AccumulationRegisterDenezhnyeSredstvaKSpisaniiu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKSpisaniius(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKSpisaniius(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKSpisaniius(Where{})
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordTypeArgs) (*AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType, error) {
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType(args.Key, nil)
}
func (r *GqlResolver) AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordTypes(ctx context.Context, args AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordTypesArgs) (*[]AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordTypes(Where{Filter: *args.Filter})
	}
	return r.Client.AccumulationRegisterDenezhnyeSredstvaKSpisaniiuRecordTypes(Where{})
}
func (r *GqlResolver) CatalogDogovoryKontragentov(ctx context.Context, args CatalogDogovoryKontragentovArgs) (*CatalogDogovoryKontragentov, error) {
	return r.Client.CatalogDogovoryKontragentov(args.Key, nil)
}
func (r *GqlResolver) CatalogDogovoryKontragentovs(ctx context.Context, args CatalogDogovoryKontragentovsArgs) (*[]CatalogDogovoryKontragentov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDogovoryKontragentovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDogovoryKontragentovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDogovoryKontragentovs(Where{})
}
func (r *GqlResolver) Order(ctx context.Context, args OrderArgs) (*Order, error) {
	return r.Client.Order(args.Key, nil)
}
func (r *GqlResolver) Orders(ctx context.Context, args OrdersArgs) (*[]Order, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Orders(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Orders(Where{Filter: *args.Filter})
	}
	return r.Client.Orders(Where{})
}
func (r *GqlResolver) DocumentChekKKMOplata(ctx context.Context, args DocumentChekKKMOplataArgs) (*DocumentChekKKMOplata, error) {
	return r.Client.DocumentChekKKMOplata(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMOplatas(ctx context.Context, args DocumentChekKKMOplatasArgs) (*[]DocumentChekKKMOplata, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMOplatas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMOplatas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMOplatas(Where{})
}
func (r *GqlResolver) DocumentChekKKMOplataSertifikatami(ctx context.Context, args DocumentChekKKMOplataSertifikatamiArgs) (*DocumentChekKKMOplataSertifikatami, error) {
	return r.Client.DocumentChekKKMOplataSertifikatami(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMOplataSertifikatamis(ctx context.Context, args DocumentChekKKMOplataSertifikatamisArgs) (*[]DocumentChekKKMOplataSertifikatami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMOplataSertifikatamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMOplataSertifikatamis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMOplataSertifikatamis(Where{})
}
func (r *GqlResolver) DocumentChekKKMProdazhaSertifikatov(ctx context.Context, args DocumentChekKKMProdazhaSertifikatovArgs) (*DocumentChekKKMProdazhaSertifikatov, error) {
	return r.Client.DocumentChekKKMProdazhaSertifikatov(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMProdazhaSertifikatovs(ctx context.Context, args DocumentChekKKMProdazhaSertifikatovsArgs) (*[]DocumentChekKKMProdazhaSertifikatov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMProdazhaSertifikatovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMProdazhaSertifikatovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMProdazhaSertifikatovs(Where{})
}
func (r *GqlResolver) DocumentChekKKMTovary(ctx context.Context, args DocumentChekKKMTovaryArgs) (*DocumentChekKKMTovary, error) {
	return r.Client.DocumentChekKKMTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMTovarys(ctx context.Context, args DocumentChekKKMTovarysArgs) (*[]DocumentChekKKMTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMTovarys(Where{})
}
func (r *GqlResolver) DocumentChekKKMDokumentyObmena(ctx context.Context, args DocumentChekKKMDokumentyObmenaArgs) (*DocumentChekKKMDokumentyObmena, error) {
	return r.Client.DocumentChekKKMDokumentyObmena(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMDokumentyObmenas(ctx context.Context, args DocumentChekKKMDokumentyObmenasArgs) (*[]DocumentChekKKMDokumentyObmena, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMDokumentyObmenas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMDokumentyObmenas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMDokumentyObmenas(Where{})
}
func (r *GqlResolver) DocumentChekKKMDogovoraRassrochkiProdazha(ctx context.Context, args DocumentChekKKMDogovoraRassrochkiProdazhaArgs) (*DocumentChekKKMDogovoraRassrochkiProdazha, error) {
	return r.Client.DocumentChekKKMDogovoraRassrochkiProdazha(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMDogovoraRassrochkiProdazhas(ctx context.Context, args DocumentChekKKMDogovoraRassrochkiProdazhasArgs) (*[]DocumentChekKKMDogovoraRassrochkiProdazha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMDogovoraRassrochkiProdazhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMDogovoraRassrochkiProdazhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMDogovoraRassrochkiProdazhas(Where{})
}
func (r *GqlResolver) DocumentChekKKMDogovoraRassrochkiOplata(ctx context.Context, args DocumentChekKKMDogovoraRassrochkiOplataArgs) (*DocumentChekKKMDogovoraRassrochkiOplata, error) {
	return r.Client.DocumentChekKKMDogovoraRassrochkiOplata(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMDogovoraRassrochkiOplatas(ctx context.Context, args DocumentChekKKMDogovoraRassrochkiOplatasArgs) (*[]DocumentChekKKMDogovoraRassrochkiOplata, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMDogovoraRassrochkiOplatas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMDogovoraRassrochkiOplatas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMDogovoraRassrochkiOplatas(Where{})
}
func (r *GqlResolver) DocumentChekKKMOplataBallami(ctx context.Context, args DocumentChekKKMOplataBallamiArgs) (*DocumentChekKKMOplataBallami, error) {
	return r.Client.DocumentChekKKMOplataBallami(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMOplataBallamis(ctx context.Context, args DocumentChekKKMOplataBallamisArgs) (*[]DocumentChekKKMOplataBallami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMOplataBallamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMOplataBallamis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMOplataBallamis(Where{})
}
func (r *GqlResolver) DocumentChekKKMSkidkiNatsenki(ctx context.Context, args DocumentChekKKMSkidkiNatsenkiArgs) (*DocumentChekKKMSkidkiNatsenki, error) {
	return r.Client.DocumentChekKKMSkidkiNatsenki(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMSkidkiNatsenkis(ctx context.Context, args DocumentChekKKMSkidkiNatsenkisArgs) (*[]DocumentChekKKMSkidkiNatsenki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMSkidkiNatsenkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMSkidkiNatsenkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMSkidkiNatsenkis(Where{})
}
func (r *GqlResolver) DocumentChekKKMUpravliaemyeSkidki(ctx context.Context, args DocumentChekKKMUpravliaemyeSkidkiArgs) (*DocumentChekKKMUpravliaemyeSkidki, error) {
	return r.Client.DocumentChekKKMUpravliaemyeSkidki(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMUpravliaemyeSkidkis(ctx context.Context, args DocumentChekKKMUpravliaemyeSkidkisArgs) (*[]DocumentChekKKMUpravliaemyeSkidki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMUpravliaemyeSkidkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMUpravliaemyeSkidkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMUpravliaemyeSkidkis(Where{})
}
func (r *GqlResolver) DocumentChekKKMPodarki(ctx context.Context, args DocumentChekKKMPodarkiArgs) (*DocumentChekKKMPodarki, error) {
	return r.Client.DocumentChekKKMPodarki(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMPodarkis(ctx context.Context, args DocumentChekKKMPodarkisArgs) (*[]DocumentChekKKMPodarki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMPodarkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMPodarkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMPodarkis(Where{})
}
func (r *GqlResolver) DocumentChekKKMKupony(ctx context.Context, args DocumentChekKKMKuponyArgs) (*DocumentChekKKMKupony, error) {
	return r.Client.DocumentChekKKMKupony(args.Key, nil)
}
func (r *GqlResolver) DocumentChekKKMKuponys(ctx context.Context, args DocumentChekKKMKuponysArgs) (*[]DocumentChekKKMKupony, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentChekKKMKuponys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentChekKKMKuponys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentChekKKMKuponys(Where{})
}
func (r *GqlResolver) DocumentPereotsenkaValiutnykhSredstv(ctx context.Context, args DocumentPereotsenkaValiutnykhSredstvArgs) (*DocumentPereotsenkaValiutnykhSredstv, error) {
	return r.Client.DocumentPereotsenkaValiutnykhSredstv(args.Key, nil)
}
func (r *GqlResolver) DocumentPereotsenkaValiutnykhSredstvs(ctx context.Context, args DocumentPereotsenkaValiutnykhSredstvsArgs) (*[]DocumentPereotsenkaValiutnykhSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPereotsenkaValiutnykhSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPereotsenkaValiutnykhSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPereotsenkaValiutnykhSredstvs(Where{})
}
func (r *GqlResolver) CatalogTipySkidokNatsenok(ctx context.Context, args CatalogTipySkidokNatsenokArgs) (*CatalogTipySkidokNatsenok, error) {
	return r.Client.CatalogTipySkidokNatsenok(args.Key, nil)
}
func (r *GqlResolver) CatalogTipySkidokNatsenoks(ctx context.Context, args CatalogTipySkidokNatsenoksArgs) (*[]CatalogTipySkidokNatsenok, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTipySkidokNatsenoks(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTipySkidokNatsenoks(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTipySkidokNatsenoks(Where{})
}
func (r *GqlResolver) CatalogTipySkidokNatsenokVremiaPoDniamNedeli(ctx context.Context, args CatalogTipySkidokNatsenokVremiaPoDniamNedeliArgs) (*CatalogTipySkidokNatsenokVremiaPoDniamNedeli, error) {
	return r.Client.CatalogTipySkidokNatsenokVremiaPoDniamNedeli(args.Key, nil)
}
func (r *GqlResolver) CatalogTipySkidokNatsenokVremiaPoDniamNedelis(ctx context.Context, args CatalogTipySkidokNatsenokVremiaPoDniamNedelisArgs) (*[]CatalogTipySkidokNatsenokVremiaPoDniamNedeli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTipySkidokNatsenokVremiaPoDniamNedelis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTipySkidokNatsenokVremiaPoDniamNedelis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTipySkidokNatsenokVremiaPoDniamNedelis(Where{})
}
func (r *GqlResolver) CatalogVidyKodirovokiTsepei(ctx context.Context, args CatalogVidyKodirovokiTsepeiArgs) (*CatalogVidyKodirovokiTsepei, error) {
	return r.Client.CatalogVidyKodirovokiTsepei(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyKodirovokiTsepeis(ctx context.Context, args CatalogVidyKodirovokiTsepeisArgs) (*[]CatalogVidyKodirovokiTsepei, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyKodirovokiTsepeis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyKodirovokiTsepeis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyKodirovokiTsepeis(Where{})
}
func (r *GqlResolver) CatalogVidyKodirovokiTsepeiElementyKodirovki(ctx context.Context, args CatalogVidyKodirovokiTsepeiElementyKodirovkiArgs) (*CatalogVidyKodirovokiTsepeiElementyKodirovki, error) {
	return r.Client.CatalogVidyKodirovokiTsepeiElementyKodirovki(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyKodirovokiTsepeiElementyKodirovkis(ctx context.Context, args CatalogVidyKodirovokiTsepeiElementyKodirovkisArgs) (*[]CatalogVidyKodirovokiTsepeiElementyKodirovki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyKodirovokiTsepeiElementyKodirovkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyKodirovokiTsepeiElementyKodirovkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyKodirovokiTsepeiElementyKodirovkis(Where{})
}
func (r *GqlResolver) CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv(ctx context.Context, args CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvArgs) (*CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv, error) {
	return r.Client.CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvs(ctx context.Context, args CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvsArgs) (*[]CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyKodirovokiTsepeiSootvetstvieZnacheniiKodrovkiZnacheniiamSvoistvs(Where{})
}
func (r *GqlResolver) DocumentOtchetKomitentuOProdazhakh(ctx context.Context, args DocumentOtchetKomitentuOProdazhakhArgs) (*DocumentOtchetKomitentuOProdazhakh, error) {
	return r.Client.DocumentOtchetKomitentuOProdazhakh(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetKomitentuOProdazhakhs(ctx context.Context, args DocumentOtchetKomitentuOProdazhakhsArgs) (*[]DocumentOtchetKomitentuOProdazhakh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetKomitentuOProdazhakhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetKomitentuOProdazhakhs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetKomitentuOProdazhakhs(Where{})
}
func (r *GqlResolver) DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva(ctx context.Context, args DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvaArgs) (*DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva, error) {
	return r.Client.DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvas(ctx context.Context, args DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvasArgs) (*[]DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstva, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetKomitentuOProdazhakhDenezhnyeSredstvas(Where{})
}
func (r *GqlResolver) DocumentOtchetKomitentuOProdazhakhTovary(ctx context.Context, args DocumentOtchetKomitentuOProdazhakhTovaryArgs) (*DocumentOtchetKomitentuOProdazhakhTovary, error) {
	return r.Client.DocumentOtchetKomitentuOProdazhakhTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetKomitentuOProdazhakhTovarys(ctx context.Context, args DocumentOtchetKomitentuOProdazhakhTovarysArgs) (*[]DocumentOtchetKomitentuOProdazhakhTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetKomitentuOProdazhakhTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetKomitentuOProdazhakhTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetKomitentuOProdazhakhTovarys(Where{})
}
func (r *GqlResolver) CatalogKassy(ctx context.Context, args CatalogKassyArgs) (*CatalogKassy, error) {
	return r.Client.CatalogKassy(args.Key, nil)
}
func (r *GqlResolver) CatalogKassys(ctx context.Context, args CatalogKassysArgs) (*[]CatalogKassy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKassys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKassys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKassys(Where{})
}
func (r *GqlResolver) CatalogKassiry(ctx context.Context, args CatalogKassiryArgs) (*CatalogKassiry, error) {
	return r.Client.CatalogKassiry(args.Key, nil)
}
func (r *GqlResolver) CatalogKassirys(ctx context.Context, args CatalogKassirysArgs) (*[]CatalogKassiry, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKassirys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKassirys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKassirys(Where{})
}
func (r *GqlResolver) DocumentZaiavkaNaPereotsenkuTovarov(ctx context.Context, args DocumentZaiavkaNaPereotsenkuTovarovArgs) (*DocumentZaiavkaNaPereotsenkuTovarov, error) {
	return r.Client.DocumentZaiavkaNaPereotsenkuTovarov(args.Key, nil)
}
func (r *GqlResolver) DocumentZaiavkaNaPereotsenkuTovarovs(ctx context.Context, args DocumentZaiavkaNaPereotsenkuTovarovsArgs) (*[]DocumentZaiavkaNaPereotsenkuTovarov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZaiavkaNaPereotsenkuTovarovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZaiavkaNaPereotsenkuTovarovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZaiavkaNaPereotsenkuTovarovs(Where{})
}
func (r *GqlResolver) DocumentZaiavkaNaPereotsenkuTovarovTovary(ctx context.Context, args DocumentZaiavkaNaPereotsenkuTovarovTovaryArgs) (*DocumentZaiavkaNaPereotsenkuTovarovTovary, error) {
	return r.Client.DocumentZaiavkaNaPereotsenkuTovarovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentZaiavkaNaPereotsenkuTovarovTovarys(ctx context.Context, args DocumentZaiavkaNaPereotsenkuTovarovTovarysArgs) (*[]DocumentZaiavkaNaPereotsenkuTovarovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZaiavkaNaPereotsenkuTovarovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZaiavkaNaPereotsenkuTovarovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZaiavkaNaPereotsenkuTovarovTovarys(Where{})
}
func (r *GqlResolver) CatalogProizvodstvennyeUchastki(ctx context.Context, args CatalogProizvodstvennyeUchastkiArgs) (*CatalogProizvodstvennyeUchastki, error) {
	return r.Client.CatalogProizvodstvennyeUchastki(args.Key, nil)
}
func (r *GqlResolver) CatalogProizvodstvennyeUchastkis(ctx context.Context, args CatalogProizvodstvennyeUchastkisArgs) (*[]CatalogProizvodstvennyeUchastki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogProizvodstvennyeUchastkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogProizvodstvennyeUchastkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogProizvodstvennyeUchastkis(Where{})
}
func (r *GqlResolver) DocumentZakrytieZakazovKlientov(ctx context.Context, args DocumentZakrytieZakazovKlientovArgs) (*DocumentZakrytieZakazovKlientov, error) {
	return r.Client.DocumentZakrytieZakazovKlientov(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytieZakazovKlientovs(ctx context.Context, args DocumentZakrytieZakazovKlientovsArgs) (*[]DocumentZakrytieZakazovKlientov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytieZakazovKlientovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytieZakazovKlientovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytieZakazovKlientovs(Where{})
}
func (r *GqlResolver) DocumentZakrytieZakazovKlientovZakazy(ctx context.Context, args DocumentZakrytieZakazovKlientovZakazyArgs) (*DocumentZakrytieZakazovKlientovZakazy, error) {
	return r.Client.DocumentZakrytieZakazovKlientovZakazy(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytieZakazovKlientovZakazys(ctx context.Context, args DocumentZakrytieZakazovKlientovZakazysArgs) (*[]DocumentZakrytieZakazovKlientovZakazy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytieZakazovKlientovZakazys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytieZakazovKlientovZakazys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytieZakazovKlientovZakazys(Where{})
}
func (r *GqlResolver) CatalogProekty(ctx context.Context, args CatalogProektyArgs) (*CatalogProekty, error) {
	return r.Client.CatalogProekty(args.Key, nil)
}
func (r *GqlResolver) CatalogProektys(ctx context.Context, args CatalogProektysArgs) (*[]CatalogProekty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogProektys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogProektys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogProektys(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieVkhodiashchee(ctx context.Context, args DocumentPlatezhnoePoruchenieVkhodiashcheeArgs) (*DocumentPlatezhnoePoruchenieVkhodiashchee, error) {
	return r.Client.DocumentPlatezhnoePoruchenieVkhodiashchee(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieVkhodiashchees(ctx context.Context, args DocumentPlatezhnoePoruchenieVkhodiashcheesArgs) (*[]DocumentPlatezhnoePoruchenieVkhodiashchee, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoePoruchenieVkhodiashchees(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoePoruchenieVkhodiashchees(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoePoruchenieVkhodiashchees(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha(ctx context.Context, args DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhaArgs) (*DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha, error) {
	return r.Client.DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhas(ctx context.Context, args DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhasArgs) (*[]DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoePoruchenieVkhodiashcheeRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta(ctx context.Context, args DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentaArgs) (*DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta, error) {
	return r.Client.DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentas(ctx context.Context, args DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentasArgs) (*[]DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoePoruchenieVkhodiashcheeRekvizityKontragentas(Where{})
}
func (r *GqlResolver) DocumentVydachaZakaza(ctx context.Context, args DocumentVydachaZakazaArgs) (*DocumentVydachaZakaza, error) {
	return r.Client.DocumentVydachaZakaza(args.Key, nil)
}
func (r *GqlResolver) DocumentVydachaZakazas(ctx context.Context, args DocumentVydachaZakazasArgs) (*[]DocumentVydachaZakaza, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVydachaZakazas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVydachaZakazas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVydachaZakazas(Where{})
}
func (r *GqlResolver) CatalogFormyOgranki(ctx context.Context, args CatalogFormyOgrankiArgs) (*CatalogFormyOgranki, error) {
	return r.Client.CatalogFormyOgranki(args.Key, nil)
}
func (r *GqlResolver) CatalogFormyOgrankis(ctx context.Context, args CatalogFormyOgrankisArgs) (*[]CatalogFormyOgranki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFormyOgrankis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFormyOgrankis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFormyOgrankis(Where{})
}
func (r *GqlResolver) CatalogFormatyMagazinov(ctx context.Context, args CatalogFormatyMagazinovArgs) (*CatalogFormatyMagazinov, error) {
	return r.Client.CatalogFormatyMagazinov(args.Key, nil)
}
func (r *GqlResolver) CatalogFormatyMagazinovs(ctx context.Context, args CatalogFormatyMagazinovsArgs) (*[]CatalogFormatyMagazinov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFormatyMagazinovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFormatyMagazinovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFormatyMagazinovs(Where{})
}
func (r *GqlResolver) CatalogRabochieMesta(ctx context.Context, args CatalogRabochieMestaArgs) (*CatalogRabochieMesta, error) {
	return r.Client.CatalogRabochieMesta(args.Key, nil)
}
func (r *GqlResolver) CatalogRabochieMestas(ctx context.Context, args CatalogRabochieMestasArgs) (*[]CatalogRabochieMesta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogRabochieMestas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogRabochieMestas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogRabochieMestas(Where{})
}
func (r *GqlResolver) CatalogNastroikiVypolneniiaObmena(ctx context.Context, args CatalogNastroikiVypolneniiaObmenaArgs) (*CatalogNastroikiVypolneniiaObmena, error) {
	return r.Client.CatalogNastroikiVypolneniiaObmena(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiVypolneniiaObmenas(ctx context.Context, args CatalogNastroikiVypolneniiaObmenasArgs) (*[]CatalogNastroikiVypolneniiaObmena, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiVypolneniiaObmenas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiVypolneniiaObmenas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiVypolneniiaObmenas(Where{})
}
func (r *GqlResolver) CatalogNastroikiVypolneniiaObmenaNastroikiObmena(ctx context.Context, args CatalogNastroikiVypolneniiaObmenaNastroikiObmenaArgs) (*CatalogNastroikiVypolneniiaObmenaNastroikiObmena, error) {
	return r.Client.CatalogNastroikiVypolneniiaObmenaNastroikiObmena(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiVypolneniiaObmenaNastroikiObmenas(ctx context.Context, args CatalogNastroikiVypolneniiaObmenaNastroikiObmenasArgs) (*[]CatalogNastroikiVypolneniiaObmenaNastroikiObmena, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiVypolneniiaObmenaNastroikiObmenas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiVypolneniiaObmenaNastroikiObmenas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiVypolneniiaObmenaNastroikiObmenas(Where{})
}
func (r *GqlResolver) CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami(ctx context.Context, args CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamiArgs) (*CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami, error) {
	return r.Client.CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamis(ctx context.Context, args CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamisArgs) (*[]CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiVypolneniiaObmenaSoobshcheniiaNeIavliaiushchiesiaOshibkamis(Where{})
}
func (r *GqlResolver) CatalogZnacheniiaSvoistvObieektov(ctx context.Context, args CatalogZnacheniiaSvoistvObieektovArgs) (*CatalogZnacheniiaSvoistvObieektov, error) {
	return r.Client.CatalogZnacheniiaSvoistvObieektov(args.Key, nil)
}
func (r *GqlResolver) CatalogZnacheniiaSvoistvObieektovs(ctx context.Context, args CatalogZnacheniiaSvoistvObieektovsArgs) (*[]CatalogZnacheniiaSvoistvObieektov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogZnacheniiaSvoistvObieektovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogZnacheniiaSvoistvObieektovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogZnacheniiaSvoistvObieektovs(Where{})
}
func (r *GqlResolver) DocumentRealizatsiiaTovarovUslug(ctx context.Context, args DocumentRealizatsiiaTovarovUslugArgs) (*DocumentRealizatsiiaTovarovUslug, error) {
	return r.Client.DocumentRealizatsiiaTovarovUslug(args.Key, nil)
}
func (r *GqlResolver) DocumentRealizatsiiaTovarovUslugs(ctx context.Context, args DocumentRealizatsiiaTovarovUslugsArgs) (*[]DocumentRealizatsiiaTovarovUslug, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRealizatsiiaTovarovUslugs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRealizatsiiaTovarovUslugs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRealizatsiiaTovarovUslugs(Where{})
}
func (r *GqlResolver) DocumentRealizatsiiaTovarovUslugTovary(ctx context.Context, args DocumentRealizatsiiaTovarovUslugTovaryArgs) (*DocumentRealizatsiiaTovarovUslugTovary, error) {
	return r.Client.DocumentRealizatsiiaTovarovUslugTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentRealizatsiiaTovarovUslugTovarys(ctx context.Context, args DocumentRealizatsiiaTovarovUslugTovarysArgs) (*[]DocumentRealizatsiiaTovarovUslugTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRealizatsiiaTovarovUslugTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRealizatsiiaTovarovUslugTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRealizatsiiaTovarovUslugTovarys(Where{})
}
func (r *GqlResolver) DocumentRealizatsiiaTovarovUslugUslugi(ctx context.Context, args DocumentRealizatsiiaTovarovUslugUslugiArgs) (*DocumentRealizatsiiaTovarovUslugUslugi, error) {
	return r.Client.DocumentRealizatsiiaTovarovUslugUslugi(args.Key, nil)
}
func (r *GqlResolver) DocumentRealizatsiiaTovarovUslugUslugis(ctx context.Context, args DocumentRealizatsiiaTovarovUslugUslugisArgs) (*[]DocumentRealizatsiiaTovarovUslugUslugi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRealizatsiiaTovarovUslugUslugis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRealizatsiiaTovarovUslugUslugis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRealizatsiiaTovarovUslugUslugis(Where{})
}
func (r *GqlResolver) DocumentSobytie(ctx context.Context, args DocumentSobytieArgs) (*DocumentSobytie, error) {
	return r.Client.DocumentSobytie(args.Key, nil)
}
func (r *GqlResolver) DocumentSobyties(ctx context.Context, args DocumentSobytiesArgs) (*[]DocumentSobytie, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSobyties(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSobyties(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSobyties(Where{})
}
func (r *GqlResolver) DocumentSobytieStoronnieLitsa(ctx context.Context, args DocumentSobytieStoronnieLitsaArgs) (*DocumentSobytieStoronnieLitsa, error) {
	return r.Client.DocumentSobytieStoronnieLitsa(args.Key, nil)
}
func (r *GqlResolver) DocumentSobytieStoronnieLitsas(ctx context.Context, args DocumentSobytieStoronnieLitsasArgs) (*[]DocumentSobytieStoronnieLitsa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSobytieStoronnieLitsas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSobytieStoronnieLitsas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSobytieStoronnieLitsas(Where{})
}
func (r *GqlResolver) CatalogVariantyOtvetovOprosov(ctx context.Context, args CatalogVariantyOtvetovOprosovArgs) (*CatalogVariantyOtvetovOprosov, error) {
	return r.Client.CatalogVariantyOtvetovOprosov(args.Key, nil)
}
func (r *GqlResolver) CatalogVariantyOtvetovOprosovs(ctx context.Context, args CatalogVariantyOtvetovOprosovsArgs) (*[]CatalogVariantyOtvetovOprosov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVariantyOtvetovOprosovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVariantyOtvetovOprosovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVariantyOtvetovOprosovs(Where{})
}
func (r *GqlResolver) CatalogGruppyPisemElektronnoiPochty(ctx context.Context, args CatalogGruppyPisemElektronnoiPochtyArgs) (*CatalogGruppyPisemElektronnoiPochty, error) {
	return r.Client.CatalogGruppyPisemElektronnoiPochty(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyPisemElektronnoiPochtys(ctx context.Context, args CatalogGruppyPisemElektronnoiPochtysArgs) (*[]CatalogGruppyPisemElektronnoiPochty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyPisemElektronnoiPochtys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyPisemElektronnoiPochtys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyPisemElektronnoiPochtys(Where{})
}
func (r *GqlResolver) CatalogGruppyPochtovoiRassylki(ctx context.Context, args CatalogGruppyPochtovoiRassylkiArgs) (*CatalogGruppyPochtovoiRassylki, error) {
	return r.Client.CatalogGruppyPochtovoiRassylki(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyPochtovoiRassylkis(ctx context.Context, args CatalogGruppyPochtovoiRassylkisArgs) (*[]CatalogGruppyPochtovoiRassylki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyPochtovoiRassylkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyPochtovoiRassylkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyPochtovoiRassylkis(Where{})
}
func (r *GqlResolver) CatalogNastroikiOtchetov(ctx context.Context, args CatalogNastroikiOtchetovArgs) (*CatalogNastroikiOtchetov, error) {
	return r.Client.CatalogNastroikiOtchetov(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiOtchetovs(ctx context.Context, args CatalogNastroikiOtchetovsArgs) (*[]CatalogNastroikiOtchetov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiOtchetovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiOtchetovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiOtchetovs(Where{})
}
func (r *GqlResolver) CatalogNastroikiOtchetovGruppyNastroekIPolzovateli(ctx context.Context, args CatalogNastroikiOtchetovGruppyNastroekIPolzovateliArgs) (*CatalogNastroikiOtchetovGruppyNastroekIPolzovateli, error) {
	return r.Client.CatalogNastroikiOtchetovGruppyNastroekIPolzovateli(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiOtchetovGruppyNastroekIPolzovatelis(ctx context.Context, args CatalogNastroikiOtchetovGruppyNastroekIPolzovatelisArgs) (*[]CatalogNastroikiOtchetovGruppyNastroekIPolzovateli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiOtchetovGruppyNastroekIPolzovatelis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiOtchetovGruppyNastroekIPolzovatelis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiOtchetovGruppyNastroekIPolzovatelis(Where{})
}
func (r *GqlResolver) CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam(ctx context.Context, args CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamArgs) (*CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam, error) {
	return r.Client.CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam(args.Key, nil)
}
func (r *GqlResolver) CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartams(ctx context.Context, args CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamsArgs) (*[]CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartams(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartams(Where{})
}
func (r *GqlResolver) CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki(ctx context.Context, args CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkiArgs) (*CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki, error) {
	return r.Client.CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki(args.Key, nil)
}
func (r *GqlResolver) CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkis(ctx context.Context, args CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkisArgs) (*[]CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSkhemyNakopitelnykhSkidokPoDiskontnymKartamSkidkis(Where{})
}
func (r *GqlResolver) Department(ctx context.Context, args DepartmentArgs) (*Department, error) {
	return r.Client.Department(args.Key, nil)
}
func (r *GqlResolver) Departments(ctx context.Context, args DepartmentsArgs) (*[]Department, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Departments(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Departments(Where{Filter: *args.Filter})
	}
	return r.Client.Departments(Where{})
}
func (r *GqlResolver) CatalogKodyVidovTovarov(ctx context.Context, args CatalogKodyVidovTovarovArgs) (*CatalogKodyVidovTovarov, error) {
	return r.Client.CatalogKodyVidovTovarov(args.Key, nil)
}
func (r *GqlResolver) CatalogKodyVidovTovarovs(ctx context.Context, args CatalogKodyVidovTovarovsArgs) (*[]CatalogKodyVidovTovarov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKodyVidovTovarovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKodyVidovTovarovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKodyVidovTovarovs(Where{})
}
func (r *GqlResolver) CatalogRassevy(ctx context.Context, args CatalogRassevyArgs) (*CatalogRassevy, error) {
	return r.Client.CatalogRassevy(args.Key, nil)
}
func (r *GqlResolver) CatalogRassevys(ctx context.Context, args CatalogRassevysArgs) (*[]CatalogRassevy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogRassevys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogRassevys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogRassevys(Where{})
}
func (r *GqlResolver) CatalogPrichinyZakrytiiaZakazov(ctx context.Context, args CatalogPrichinyZakrytiiaZakazovArgs) (*CatalogPrichinyZakrytiiaZakazov, error) {
	return r.Client.CatalogPrichinyZakrytiiaZakazov(args.Key, nil)
}
func (r *GqlResolver) CatalogPrichinyZakrytiiaZakazovs(ctx context.Context, args CatalogPrichinyZakrytiiaZakazovsArgs) (*[]CatalogPrichinyZakrytiiaZakazov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogPrichinyZakrytiiaZakazovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogPrichinyZakrytiiaZakazovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogPrichinyZakrytiiaZakazovs(Where{})
}
func (r *GqlResolver) CatalogSegmentyNomenklatury(ctx context.Context, args CatalogSegmentyNomenklaturyArgs) (*CatalogSegmentyNomenklatury, error) {
	return r.Client.CatalogSegmentyNomenklatury(args.Key, nil)
}
func (r *GqlResolver) CatalogSegmentyNomenklaturys(ctx context.Context, args CatalogSegmentyNomenklaturysArgs) (*[]CatalogSegmentyNomenklatury, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSegmentyNomenklaturys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSegmentyNomenklaturys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSegmentyNomenklaturys(Where{})
}
func (r *GqlResolver) CatalogSostavStrokiCheka(ctx context.Context, args CatalogSostavStrokiChekaArgs) (*CatalogSostavStrokiCheka, error) {
	return r.Client.CatalogSostavStrokiCheka(args.Key, nil)
}
func (r *GqlResolver) CatalogSostavStrokiChekas(ctx context.Context, args CatalogSostavStrokiChekasArgs) (*[]CatalogSostavStrokiCheka, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSostavStrokiChekas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSostavStrokiChekas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSostavStrokiChekas(Where{})
}
func (r *GqlResolver) CatalogUsloviiaPriemaIzdeliiNaKomissiiu(ctx context.Context, args CatalogUsloviiaPriemaIzdeliiNaKomissiiuArgs) (*CatalogUsloviiaPriemaIzdeliiNaKomissiiu, error) {
	return r.Client.CatalogUsloviiaPriemaIzdeliiNaKomissiiu(args.Key, nil)
}
func (r *GqlResolver) CatalogUsloviiaPriemaIzdeliiNaKomissiius(ctx context.Context, args CatalogUsloviiaPriemaIzdeliiNaKomissiiusArgs) (*[]CatalogUsloviiaPriemaIzdeliiNaKomissiiu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUsloviiaPriemaIzdeliiNaKomissiius(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUsloviiaPriemaIzdeliiNaKomissiius(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUsloviiaPriemaIzdeliiNaKomissiius(Where{})
}
func (r *GqlResolver) CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok(ctx context.Context, args CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenokArgs) (*CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok, error) {
	return r.Client.CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok(args.Key, nil)
}
func (r *GqlResolver) CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenoks(ctx context.Context, args CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenoksArgs) (*[]CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenok, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenoks(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenoks(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUsloviiaPriemaIzdeliiNaKomissiiuGrafikUtsenoks(Where{})
}
func (r *GqlResolver) CatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei(ctx context.Context, args CatalogIstochnikiInformatsiiPriObrashcheniiPokupateleiArgs) (*CatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei, error) {
	return r.Client.CatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei(args.Key, nil)
}
func (r *GqlResolver) CatalogIstochnikiInformatsiiPriObrashcheniiPokupateleis(ctx context.Context, args CatalogIstochnikiInformatsiiPriObrashcheniiPokupateleisArgs) (*[]CatalogIstochnikiInformatsiiPriObrashcheniiPokupatelei, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogIstochnikiInformatsiiPriObrashcheniiPokupateleis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogIstochnikiInformatsiiPriObrashcheniiPokupateleis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogIstochnikiInformatsiiPriObrashcheniiPokupateleis(Where{})
}
func (r *GqlResolver) DocumentKorrektirovkaDolga(ctx context.Context, args DocumentKorrektirovkaDolgaArgs) (*DocumentKorrektirovkaDolga, error) {
	return r.Client.DocumentKorrektirovkaDolga(args.Key, nil)
}
func (r *GqlResolver) DocumentKorrektirovkaDolgas(ctx context.Context, args DocumentKorrektirovkaDolgasArgs) (*[]DocumentKorrektirovkaDolga, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKorrektirovkaDolgas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKorrektirovkaDolgas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKorrektirovkaDolgas(Where{})
}
func (r *GqlResolver) DocumentKorrektirovkaDolgaSummyDolga(ctx context.Context, args DocumentKorrektirovkaDolgaSummyDolgaArgs) (*DocumentKorrektirovkaDolgaSummyDolga, error) {
	return r.Client.DocumentKorrektirovkaDolgaSummyDolga(args.Key, nil)
}
func (r *GqlResolver) DocumentKorrektirovkaDolgaSummyDolgas(ctx context.Context, args DocumentKorrektirovkaDolgaSummyDolgasArgs) (*[]DocumentKorrektirovkaDolgaSummyDolga, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKorrektirovkaDolgaSummyDolgas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKorrektirovkaDolgaSummyDolgas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKorrektirovkaDolgaSummyDolgas(Where{})
}
func (r *GqlResolver) PayType(ctx context.Context, args PayTypeArgs) (*PayType, error) {
	return r.Client.PayType(args.Key, nil)
}
func (r *GqlResolver) PayTypes(ctx context.Context, args PayTypesArgs) (*[]PayType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.PayTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.PayTypes(Where{Filter: *args.Filter})
	}
	return r.Client.PayTypes(Where{})
}
func (r *GqlResolver) CatalogKhranilishcheShablonov(ctx context.Context, args CatalogKhranilishcheShablonovArgs) (*CatalogKhranilishcheShablonov, error) {
	return r.Client.CatalogKhranilishcheShablonov(args.Key, nil)
}
func (r *GqlResolver) CatalogKhranilishcheShablonovs(ctx context.Context, args CatalogKhranilishcheShablonovsArgs) (*[]CatalogKhranilishcheShablonov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKhranilishcheShablonovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKhranilishcheShablonovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKhranilishcheShablonovs(Where{})
}
func (r *GqlResolver) DocumentZaiavkaNaRaskhodovanieSredstv(ctx context.Context, args DocumentZaiavkaNaRaskhodovanieSredstvArgs) (*DocumentZaiavkaNaRaskhodovanieSredstv, error) {
	return r.Client.DocumentZaiavkaNaRaskhodovanieSredstv(args.Key, nil)
}
func (r *GqlResolver) DocumentZaiavkaNaRaskhodovanieSredstvs(ctx context.Context, args DocumentZaiavkaNaRaskhodovanieSredstvsArgs) (*[]DocumentZaiavkaNaRaskhodovanieSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvs(Where{})
}
func (r *GqlResolver) DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha(ctx context.Context, args DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhaArgs) (*DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha, error) {
	return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhas(ctx context.Context, args DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhasArgs) (*[]DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki(ctx context.Context, args DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkiArgs) (*DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki, error) {
	return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki(args.Key, nil)
}
func (r *GqlResolver) DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkis(ctx context.Context, args DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkisArgs) (*[]DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZaiavkaNaRaskhodovanieSredstvRazmeshchenieZaiavkis(Where{})
}
func (r *GqlResolver) DocumentZakrytieZakazovPostavshchikam(ctx context.Context, args DocumentZakrytieZakazovPostavshchikamArgs) (*DocumentZakrytieZakazovPostavshchikam, error) {
	return r.Client.DocumentZakrytieZakazovPostavshchikam(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytieZakazovPostavshchikams(ctx context.Context, args DocumentZakrytieZakazovPostavshchikamsArgs) (*[]DocumentZakrytieZakazovPostavshchikam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytieZakazovPostavshchikams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytieZakazovPostavshchikams(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytieZakazovPostavshchikams(Where{})
}
func (r *GqlResolver) DocumentZakrytieZakazovPostavshchikamZakazy(ctx context.Context, args DocumentZakrytieZakazovPostavshchikamZakazyArgs) (*DocumentZakrytieZakazovPostavshchikamZakazy, error) {
	return r.Client.DocumentZakrytieZakazovPostavshchikamZakazy(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytieZakazovPostavshchikamZakazys(ctx context.Context, args DocumentZakrytieZakazovPostavshchikamZakazysArgs) (*[]DocumentZakrytieZakazovPostavshchikamZakazy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytieZakazovPostavshchikamZakazys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytieZakazovPostavshchikamZakazys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytieZakazovPostavshchikamZakazys(Where{})
}
func (r *GqlResolver) CatalogVidyKamnei(ctx context.Context, args CatalogVidyKamneiArgs) (*CatalogVidyKamnei, error) {
	return r.Client.CatalogVidyKamnei(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyKamneis(ctx context.Context, args CatalogVidyKamneisArgs) (*[]CatalogVidyKamnei, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyKamneis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyKamneis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyKamneis(Where{})
}
func (r *GqlResolver) DocumentAnketyKlientovDliaFinMonitoringa(ctx context.Context, args DocumentAnketyKlientovDliaFinMonitoringaArgs) (*DocumentAnketyKlientovDliaFinMonitoringa, error) {
	return r.Client.DocumentAnketyKlientovDliaFinMonitoringa(args.Key, nil)
}
func (r *GqlResolver) DocumentAnketyKlientovDliaFinMonitoringas(ctx context.Context, args DocumentAnketyKlientovDliaFinMonitoringasArgs) (*[]DocumentAnketyKlientovDliaFinMonitoringa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAnketyKlientovDliaFinMonitoringas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAnketyKlientovDliaFinMonitoringas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAnketyKlientovDliaFinMonitoringas(Where{})
}
func (r *GqlResolver) DocumentAnketyKlientovDliaFinMonitoringaAnkety(ctx context.Context, args DocumentAnketyKlientovDliaFinMonitoringaAnketyArgs) (*DocumentAnketyKlientovDliaFinMonitoringaAnkety, error) {
	return r.Client.DocumentAnketyKlientovDliaFinMonitoringaAnkety(args.Key, nil)
}
func (r *GqlResolver) DocumentAnketyKlientovDliaFinMonitoringaAnketys(ctx context.Context, args DocumentAnketyKlientovDliaFinMonitoringaAnketysArgs) (*[]DocumentAnketyKlientovDliaFinMonitoringaAnkety, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAnketyKlientovDliaFinMonitoringaAnketys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAnketyKlientovDliaFinMonitoringaAnketys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAnketyKlientovDliaFinMonitoringaAnketys(Where{})
}
func (r *GqlResolver) CatalogDogovoryRassrochki(ctx context.Context, args CatalogDogovoryRassrochkiArgs) (*CatalogDogovoryRassrochki, error) {
	return r.Client.CatalogDogovoryRassrochki(args.Key, nil)
}
func (r *GqlResolver) CatalogDogovoryRassrochkis(ctx context.Context, args CatalogDogovoryRassrochkisArgs) (*[]CatalogDogovoryRassrochki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDogovoryRassrochkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDogovoryRassrochkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDogovoryRassrochkis(Where{})
}
func (r *GqlResolver) CatalogSertifikaty(ctx context.Context, args CatalogSertifikatyArgs) (*CatalogSertifikaty, error) {
	return r.Client.CatalogSertifikaty(args.Key, nil)
}
func (r *GqlResolver) CatalogSertifikatys(ctx context.Context, args CatalogSertifikatysArgs) (*[]CatalogSertifikaty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSertifikatys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSertifikatys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSertifikatys(Where{})
}
func (r *GqlResolver) DocumentPostuplenieDavalcheskogoMetalla(ctx context.Context, args DocumentPostuplenieDavalcheskogoMetallaArgs) (*DocumentPostuplenieDavalcheskogoMetalla, error) {
	return r.Client.DocumentPostuplenieDavalcheskogoMetalla(args.Key, nil)
}
func (r *GqlResolver) DocumentPostuplenieDavalcheskogoMetallas(ctx context.Context, args DocumentPostuplenieDavalcheskogoMetallasArgs) (*[]DocumentPostuplenieDavalcheskogoMetalla, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPostuplenieDavalcheskogoMetallas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPostuplenieDavalcheskogoMetallas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPostuplenieDavalcheskogoMetallas(Where{})
}
func (r *GqlResolver) DocumentInkassovoePorucheniePeredannoe(ctx context.Context, args DocumentInkassovoePorucheniePeredannoeArgs) (*DocumentInkassovoePorucheniePeredannoe, error) {
	return r.Client.DocumentInkassovoePorucheniePeredannoe(args.Key, nil)
}
func (r *GqlResolver) DocumentInkassovoePorucheniePeredannoes(ctx context.Context, args DocumentInkassovoePorucheniePeredannoesArgs) (*[]DocumentInkassovoePorucheniePeredannoe, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInkassovoePorucheniePeredannoes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInkassovoePorucheniePeredannoes(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInkassovoePorucheniePeredannoes(Where{})
}
func (r *GqlResolver) DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha(ctx context.Context, args DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhaArgs) (*DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha, error) {
	return r.Client.DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhas(ctx context.Context, args DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhasArgs) (*[]DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInkassovoePorucheniePeredannoeRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentInkassovoePorucheniePeredannoeRekvizityKontragenta(ctx context.Context, args DocumentInkassovoePorucheniePeredannoeRekvizityKontragentaArgs) (*DocumentInkassovoePorucheniePeredannoeRekvizityKontragenta, error) {
	return r.Client.DocumentInkassovoePorucheniePeredannoeRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentInkassovoePorucheniePeredannoeRekvizityKontragentas(ctx context.Context, args DocumentInkassovoePorucheniePeredannoeRekvizityKontragentasArgs) (*[]DocumentInkassovoePorucheniePeredannoeRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInkassovoePorucheniePeredannoeRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInkassovoePorucheniePeredannoeRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInkassovoePorucheniePeredannoeRekvizityKontragentas(Where{})
}
func (r *GqlResolver) CatalogFormulyDliaRascheta(ctx context.Context, args CatalogFormulyDliaRaschetaArgs) (*CatalogFormulyDliaRascheta, error) {
	return r.Client.CatalogFormulyDliaRascheta(args.Key, nil)
}
func (r *GqlResolver) CatalogFormulyDliaRaschetas(ctx context.Context, args CatalogFormulyDliaRaschetasArgs) (*[]CatalogFormulyDliaRascheta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFormulyDliaRaschetas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFormulyDliaRaschetas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFormulyDliaRaschetas(Where{})
}
func (r *GqlResolver) CatalogKupony(ctx context.Context, args CatalogKuponyArgs) (*CatalogKupony, error) {
	return r.Client.CatalogKupony(args.Key, nil)
}
func (r *GqlResolver) CatalogKuponys(ctx context.Context, args CatalogKuponysArgs) (*[]CatalogKupony, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKuponys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKuponys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKuponys(Where{})
}
func (r *GqlResolver) Correcting(ctx context.Context, args CorrectingArgs) (*Correcting, error) {
	return r.Client.Correcting(args.Key, nil)
}
func (r *GqlResolver) Correctings(ctx context.Context, args CorrectingsArgs) (*[]Correcting, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Correctings(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Correctings(Where{Filter: *args.Filter})
	}
	return r.Client.Correctings(Where{})
}
func (r *GqlResolver) DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia(ctx context.Context, args DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniiaArgs) (*DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia, error) {
	return r.Client.DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia(args.Key, nil)
}
func (r *GqlResolver) DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniias(ctx context.Context, args DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniiasArgs) (*[]DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniias(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKorrektirovkaZapiseiRegistrovNakopleniiaTablitsaRegistrovNakopleniias(Where{})
}
func (r *GqlResolver) DocumentInternetZakaz(ctx context.Context, args DocumentInternetZakazArgs) (*DocumentInternetZakaz, error) {
	return r.Client.DocumentInternetZakaz(args.Key, nil)
}
func (r *GqlResolver) DocumentInternetZakazs(ctx context.Context, args DocumentInternetZakazsArgs) (*[]DocumentInternetZakaz, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInternetZakazs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInternetZakazs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInternetZakazs(Where{})
}
func (r *GqlResolver) DocumentInternetZakazTovaryInternetZakaza(ctx context.Context, args DocumentInternetZakazTovaryInternetZakazaArgs) (*DocumentInternetZakazTovaryInternetZakaza, error) {
	return r.Client.DocumentInternetZakazTovaryInternetZakaza(args.Key, nil)
}
func (r *GqlResolver) DocumentInternetZakazTovaryInternetZakazas(ctx context.Context, args DocumentInternetZakazTovaryInternetZakazasArgs) (*[]DocumentInternetZakazTovaryInternetZakaza, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInternetZakazTovaryInternetZakazas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInternetZakazTovaryInternetZakazas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInternetZakazTovaryInternetZakazas(Where{})
}
func (r *GqlResolver) DocumentInternetZakazTovary(ctx context.Context, args DocumentInternetZakazTovaryArgs) (*DocumentInternetZakazTovary, error) {
	return r.Client.DocumentInternetZakazTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentInternetZakazTovarys(ctx context.Context, args DocumentInternetZakazTovarysArgs) (*[]DocumentInternetZakazTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInternetZakazTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInternetZakazTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInternetZakazTovarys(Where{})
}
func (r *GqlResolver) CatalogRegiony(ctx context.Context, args CatalogRegionyArgs) (*CatalogRegiony, error) {
	return r.Client.CatalogRegiony(args.Key, nil)
}
func (r *GqlResolver) CatalogRegionys(ctx context.Context, args CatalogRegionysArgs) (*[]CatalogRegiony, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogRegionys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogRegionys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogRegionys(Where{})
}
func (r *GqlResolver) SaleJournal(ctx context.Context, args SaleJournalArgs) (*SaleJournal, error) {
	return r.Client.SaleJournal(args.Key, nil)
}
func (r *GqlResolver) SaleJournals(ctx context.Context, args SaleJournalsArgs) (*[]SaleJournal, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.SaleJournals(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.SaleJournals(Where{Filter: *args.Filter})
	}
	return r.Client.SaleJournals(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhBonusy(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhBonusyArgs) (*DocumentOtchetORoznichnykhProdazhakhBonusy, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhBonusy(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhBonusys(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhBonusysArgs) (*[]DocumentOtchetORoznichnykhProdazhakhBonusy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhBonusys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhBonusys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhBonusys(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamiArgs) (*DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamis(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamisArgs) (*[]DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataBankovskimiKreditamis(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamiArgs) (*DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamis(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamisArgs) (*[]DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataPlatezhnymiKartamis(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhOplataSertifikatami(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamiArgs) (*DocumentOtchetORoznichnykhProdazhakhOplataSertifikatami, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataSertifikatami(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamis(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamisArgs) (*[]DocumentOtchetORoznichnykhProdazhakhOplataSertifikatami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataSertifikatamis(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovArgs) (*DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovs(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovsArgs) (*[]DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhProdazhaSertifikatovs(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhTovary(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhTovaryArgs) (*DocumentOtchetORoznichnykhProdazhakhTovary, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhTovarys(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhTovarysArgs) (*[]DocumentOtchetORoznichnykhProdazhakhTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhTovarys(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhaArgs) (*DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhas(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhasArgs) (*[]DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiProdazhas(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhDokumentyObmena(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhDokumentyObmenaArgs) (*DocumentOtchetORoznichnykhProdazhakhDokumentyObmena, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhDokumentyObmena(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhDokumentyObmenas(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhDokumentyObmenasArgs) (*[]DocumentOtchetORoznichnykhProdazhakhDokumentyObmena, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhDokumentyObmenas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhDokumentyObmenas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhDokumentyObmenas(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplataArgs) (*DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplatas(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplatasArgs) (*[]DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplata, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplatas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplatas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhDogovoraRassrochkiOplatas(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhOplataBallami(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhOplataBallamiArgs) (*DocumentOtchetORoznichnykhProdazhakhOplataBallami, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataBallami(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhOplataBallamis(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhOplataBallamisArgs) (*[]DocumentOtchetORoznichnykhProdazhakhOplataBallami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataBallamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataBallamis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhOplataBallamis(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkiArgs) (*DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkis(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkisArgs) (*[]DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhSkidkiNatsenkis(Where{})
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhKupony(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhKuponyArgs) (*DocumentOtchetORoznichnykhProdazhakhKupony, error) {
	return r.Client.DocumentOtchetORoznichnykhProdazhakhKupony(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetORoznichnykhProdazhakhKuponys(ctx context.Context, args DocumentOtchetORoznichnykhProdazhakhKuponysArgs) (*[]DocumentOtchetORoznichnykhProdazhakhKupony, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetORoznichnykhProdazhakhKuponys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetORoznichnykhProdazhakhKuponys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetORoznichnykhProdazhakhKuponys(Where{})
}
func (r *GqlResolver) DocumentOtmenaSkidokNomenklatury(ctx context.Context, args DocumentOtmenaSkidokNomenklaturyArgs) (*DocumentOtmenaSkidokNomenklatury, error) {
	return r.Client.DocumentOtmenaSkidokNomenklatury(args.Key, nil)
}
func (r *GqlResolver) DocumentOtmenaSkidokNomenklaturys(ctx context.Context, args DocumentOtmenaSkidokNomenklaturysArgs) (*[]DocumentOtmenaSkidokNomenklatury, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtmenaSkidokNomenklaturys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtmenaSkidokNomenklaturys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtmenaSkidokNomenklaturys(Where{})
}
func (r *GqlResolver) DocumentOtmenaSkidokNomenklaturyDokumenty(ctx context.Context, args DocumentOtmenaSkidokNomenklaturyDokumentyArgs) (*DocumentOtmenaSkidokNomenklaturyDokumenty, error) {
	return r.Client.DocumentOtmenaSkidokNomenklaturyDokumenty(args.Key, nil)
}
func (r *GqlResolver) DocumentOtmenaSkidokNomenklaturyDokumentys(ctx context.Context, args DocumentOtmenaSkidokNomenklaturyDokumentysArgs) (*[]DocumentOtmenaSkidokNomenklaturyDokumenty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtmenaSkidokNomenklaturyDokumentys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtmenaSkidokNomenklaturyDokumentys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtmenaSkidokNomenklaturyDokumentys(Where{})
}
func (r *GqlResolver) CatalogTovarnyeGruppy(ctx context.Context, args CatalogTovarnyeGruppyArgs) (*CatalogTovarnyeGruppy, error) {
	return r.Client.CatalogTovarnyeGruppy(args.Key, nil)
}
func (r *GqlResolver) CatalogTovarnyeGruppys(ctx context.Context, args CatalogTovarnyeGruppysArgs) (*[]CatalogTovarnyeGruppy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTovarnyeGruppys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTovarnyeGruppys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTovarnyeGruppys(Where{})
}
func (r *GqlResolver) DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv(ctx context.Context, args DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvArgs) (*DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv, error) {
	return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvs(ctx context.Context, args DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvsArgs) (*[]DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvs(Where{})
}
func (r *GqlResolver) DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha(ctx context.Context, args DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaArgs) (*DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha, error) {
	return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(ctx context.Context, args DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhasArgs) (*[]DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta(ctx context.Context, args DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentaArgs) (*DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta, error) {
	return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentas(ctx context.Context, args DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentasArgs) (*[]DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnyiOrderPostuplenieDenezhnykhSredstvRekvizityKontragentas(Where{})
}
func (r *GqlResolver) CatalogOrderKey(ctx context.Context, args CatalogOrderKeyArgs) (*CatalogOrderKey, error) {
	return r.Client.CatalogOrderKey(args.Key, nil)
}
func (r *GqlResolver) CatalogOrderKeys(ctx context.Context, args CatalogOrderKeysArgs) (*[]CatalogOrderKey, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogOrderKeys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogOrderKeys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogOrderKeys(Where{})
}
func (r *GqlResolver) DocumentKassovyiChekKorrektsii(ctx context.Context, args DocumentKassovyiChekKorrektsiiArgs) (*DocumentKassovyiChekKorrektsii, error) {
	return r.Client.DocumentKassovyiChekKorrektsii(args.Key, nil)
}
func (r *GqlResolver) DocumentKassovyiChekKorrektsiis(ctx context.Context, args DocumentKassovyiChekKorrektsiisArgs) (*[]DocumentKassovyiChekKorrektsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKassovyiChekKorrektsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKassovyiChekKorrektsiis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKassovyiChekKorrektsiis(Where{})
}
func (r *GqlResolver) DocumentKassovyiChekKorrektsiiOplata(ctx context.Context, args DocumentKassovyiChekKorrektsiiOplataArgs) (*DocumentKassovyiChekKorrektsiiOplata, error) {
	return r.Client.DocumentKassovyiChekKorrektsiiOplata(args.Key, nil)
}
func (r *GqlResolver) DocumentKassovyiChekKorrektsiiOplatas(ctx context.Context, args DocumentKassovyiChekKorrektsiiOplatasArgs) (*[]DocumentKassovyiChekKorrektsiiOplata, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKassovyiChekKorrektsiiOplatas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKassovyiChekKorrektsiiOplatas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKassovyiChekKorrektsiiOplatas(Where{})
}
func (r *GqlResolver) DocumentSchetNaOplatuPokupateliu(ctx context.Context, args DocumentSchetNaOplatuPokupateliuArgs) (*DocumentSchetNaOplatuPokupateliu, error) {
	return r.Client.DocumentSchetNaOplatuPokupateliu(args.Key, nil)
}
func (r *GqlResolver) DocumentSchetNaOplatuPokupatelius(ctx context.Context, args DocumentSchetNaOplatuPokupateliusArgs) (*[]DocumentSchetNaOplatuPokupateliu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSchetNaOplatuPokupatelius(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSchetNaOplatuPokupatelius(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSchetNaOplatuPokupatelius(Where{})
}
func (r *GqlResolver) DocumentSchetNaOplatuPokupateliuTovary(ctx context.Context, args DocumentSchetNaOplatuPokupateliuTovaryArgs) (*DocumentSchetNaOplatuPokupateliuTovary, error) {
	return r.Client.DocumentSchetNaOplatuPokupateliuTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentSchetNaOplatuPokupateliuTovarys(ctx context.Context, args DocumentSchetNaOplatuPokupateliuTovarysArgs) (*[]DocumentSchetNaOplatuPokupateliuTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSchetNaOplatuPokupateliuTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSchetNaOplatuPokupateliuTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSchetNaOplatuPokupateliuTovarys(Where{})
}
func (r *GqlResolver) DocumentSchetNaOplatuPokupateliuUslugi(ctx context.Context, args DocumentSchetNaOplatuPokupateliuUslugiArgs) (*DocumentSchetNaOplatuPokupateliuUslugi, error) {
	return r.Client.DocumentSchetNaOplatuPokupateliuUslugi(args.Key, nil)
}
func (r *GqlResolver) DocumentSchetNaOplatuPokupateliuUslugis(ctx context.Context, args DocumentSchetNaOplatuPokupateliuUslugisArgs) (*[]DocumentSchetNaOplatuPokupateliuUslugi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSchetNaOplatuPokupateliuUslugis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSchetNaOplatuPokupateliuUslugis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSchetNaOplatuPokupateliuUslugis(Where{})
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymi(ctx context.Context, args CatalogNastroikiObmenaDannymiArgs) (*CatalogNastroikiObmenaDannymi, error) {
	return r.Client.CatalogNastroikiObmenaDannymi(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymis(ctx context.Context, args CatalogNastroikiObmenaDannymisArgs) (*[]CatalogNastroikiObmenaDannymi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiObmenaDannymis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiObmenaDannymis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiObmenaDannymis(Where{})
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov(ctx context.Context, args CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovArgs) (*CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov, error) {
	return r.Client.CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovs(ctx context.Context, args CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovsArgs) (*[]CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiObmenaDannymiNastroikaVariantovPoiskaObieektovs(Where{})
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh(ctx context.Context, args CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhArgs) (*CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh, error) {
	return r.Client.CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhs(ctx context.Context, args CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhsArgs) (*[]CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiObmenaDannymiNastroikaVygruzkiDannykhs(Where{})
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami(ctx context.Context, args CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamiArgs) (*CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami, error) {
	return r.Client.CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamis(ctx context.Context, args CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamisArgs) (*[]CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiObmenaDannymiSoobshcheniiaNeIavliaiushchiesiaOshibkamis(Where{})
}
func (r *GqlResolver) DocumentJournalBankovskieRaschetnyeDokumenty(ctx context.Context, args DocumentJournalBankovskieRaschetnyeDokumentyArgs) (*DocumentJournalBankovskieRaschetnyeDokumenty, error) {
	return r.Client.DocumentJournalBankovskieRaschetnyeDokumenty(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalBankovskieRaschetnyeDokumentys(ctx context.Context, args DocumentJournalBankovskieRaschetnyeDokumentysArgs) (*[]DocumentJournalBankovskieRaschetnyeDokumenty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalBankovskieRaschetnyeDokumentys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalBankovskieRaschetnyeDokumentys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalBankovskieRaschetnyeDokumentys(Where{})
}
func (r *GqlResolver) DocumentZamenaDiskontnoiKarty(ctx context.Context, args DocumentZamenaDiskontnoiKartyArgs) (*DocumentZamenaDiskontnoiKarty, error) {
	return r.Client.DocumentZamenaDiskontnoiKarty(args.Key, nil)
}
func (r *GqlResolver) DocumentZamenaDiskontnoiKartys(ctx context.Context, args DocumentZamenaDiskontnoiKartysArgs) (*[]DocumentZamenaDiskontnoiKarty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZamenaDiskontnoiKartys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZamenaDiskontnoiKartys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZamenaDiskontnoiKartys(Where{})
}
func (r *GqlResolver) ReturnToSupplier(ctx context.Context, args ReturnToSupplierArgs) (*ReturnToSupplier, error) {
	return r.Client.ReturnToSupplier(args.Key, nil)
}
func (r *GqlResolver) ReturnToSuppliers(ctx context.Context, args ReturnToSuppliersArgs) (*[]ReturnToSupplier, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.ReturnToSuppliers(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.ReturnToSuppliers(Where{Filter: *args.Filter})
	}
	return r.Client.ReturnToSuppliers(Where{})
}
func (r *GqlResolver) DocumentVozvratTovarovPostavshchikuTovary(ctx context.Context, args DocumentVozvratTovarovPostavshchikuTovaryArgs) (*DocumentVozvratTovarovPostavshchikuTovary, error) {
	return r.Client.DocumentVozvratTovarovPostavshchikuTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratTovarovPostavshchikuTovarys(ctx context.Context, args DocumentVozvratTovarovPostavshchikuTovarysArgs) (*[]DocumentVozvratTovarovPostavshchikuTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratTovarovPostavshchikuTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratTovarovPostavshchikuTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratTovarovPostavshchikuTovarys(Where{})
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSklade(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladeArgs) (*DocumentInventarizatsiiaTovarovNaSklade, error) {
	return r.Client.DocumentInventarizatsiiaTovarovNaSklade(args.Key, nil)
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSklades(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladesArgs) (*[]DocumentInventarizatsiiaTovarovNaSklade, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInventarizatsiiaTovarovNaSklades(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInventarizatsiiaTovarovNaSklades(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInventarizatsiiaTovarovNaSklades(Where{})
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSkladeTovary(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladeTovaryArgs) (*DocumentInventarizatsiiaTovarovNaSkladeTovary, error) {
	return r.Client.DocumentInventarizatsiiaTovarovNaSkladeTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSkladeTovarys(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladeTovarysArgs) (*[]DocumentInventarizatsiiaTovarovNaSkladeTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInventarizatsiiaTovarovNaSkladeTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInventarizatsiiaTovarovNaSkladeTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInventarizatsiiaTovarovNaSkladeTovarys(Where{})
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiArgs) (*DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii, error) {
	return r.Client.DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii(args.Key, nil)
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiisArgs) (*[]DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(Where{})
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSkladeSertifikaty(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladeSertifikatyArgs) (*DocumentInventarizatsiiaTovarovNaSkladeSertifikaty, error) {
	return r.Client.DocumentInventarizatsiiaTovarovNaSkladeSertifikaty(args.Key, nil)
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSkladeSertifikatys(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladeSertifikatysArgs) (*[]DocumentInventarizatsiiaTovarovNaSkladeSertifikaty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInventarizatsiiaTovarovNaSkladeSertifikatys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInventarizatsiiaTovarovNaSkladeSertifikatys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInventarizatsiiaTovarovNaSkladeSertifikatys(Where{})
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutiArgs) (*DocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti, error) {
	return r.Client.DocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti(args.Key, nil)
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutis(ctx context.Context, args DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutisArgs) (*[]DocumentInventarizatsiiaTovarovNaSkladeTovaryVPuti, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInventarizatsiiaTovarovNaSkladeTovaryVPutis(Where{})
}
func (r *GqlResolver) DocumentPrikhodnyiKassovyiOrder(ctx context.Context, args DocumentPrikhodnyiKassovyiOrderArgs) (*DocumentPrikhodnyiKassovyiOrder, error) {
	return r.Client.DocumentPrikhodnyiKassovyiOrder(args.Key, nil)
}
func (r *GqlResolver) DocumentPrikhodnyiKassovyiOrders(ctx context.Context, args DocumentPrikhodnyiKassovyiOrdersArgs) (*[]DocumentPrikhodnyiKassovyiOrder, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPrikhodnyiKassovyiOrders(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPrikhodnyiKassovyiOrders(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPrikhodnyiKassovyiOrders(Where{})
}
func (r *GqlResolver) DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha(ctx context.Context, args DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhaArgs) (*DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha, error) {
	return r.Client.DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhas(ctx context.Context, args DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhasArgs) (*[]DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPrikhodnyiKassovyiOrderRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentPrikhodnyiKassovyiOrderOplata(ctx context.Context, args DocumentPrikhodnyiKassovyiOrderOplataArgs) (*DocumentPrikhodnyiKassovyiOrderOplata, error) {
	return r.Client.DocumentPrikhodnyiKassovyiOrderOplata(args.Key, nil)
}
func (r *GqlResolver) DocumentPrikhodnyiKassovyiOrderOplatas(ctx context.Context, args DocumentPrikhodnyiKassovyiOrderOplatasArgs) (*[]DocumentPrikhodnyiKassovyiOrderOplata, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPrikhodnyiKassovyiOrderOplatas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPrikhodnyiKassovyiOrderOplatas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPrikhodnyiKassovyiOrderOplatas(Where{})
}
func (r *GqlResolver) DocumentPrikhodnyiKassovyiOrderTovary(ctx context.Context, args DocumentPrikhodnyiKassovyiOrderTovaryArgs) (*DocumentPrikhodnyiKassovyiOrderTovary, error) {
	return r.Client.DocumentPrikhodnyiKassovyiOrderTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPrikhodnyiKassovyiOrderTovarys(ctx context.Context, args DocumentPrikhodnyiKassovyiOrderTovarysArgs) (*[]DocumentPrikhodnyiKassovyiOrderTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPrikhodnyiKassovyiOrderTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPrikhodnyiKassovyiOrderTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPrikhodnyiKassovyiOrderTovarys(Where{})
}
func (r *GqlResolver) CatalogPrichinyVozvrata(ctx context.Context, args CatalogPrichinyVozvrataArgs) (*CatalogPrichinyVozvrata, error) {
	return r.Client.CatalogPrichinyVozvrata(args.Key, nil)
}
func (r *GqlResolver) CatalogPrichinyVozvratas(ctx context.Context, args CatalogPrichinyVozvratasArgs) (*[]CatalogPrichinyVozvrata, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogPrichinyVozvratas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogPrichinyVozvratas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogPrichinyVozvratas(Where{})
}
func (r *GqlResolver) DocumentDenezhnyiChek(ctx context.Context, args DocumentDenezhnyiChekArgs) (*DocumentDenezhnyiChek, error) {
	return r.Client.DocumentDenezhnyiChek(args.Key, nil)
}
func (r *GqlResolver) DocumentDenezhnyiCheks(ctx context.Context, args DocumentDenezhnyiCheksArgs) (*[]DocumentDenezhnyiChek, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentDenezhnyiCheks(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentDenezhnyiCheks(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentDenezhnyiCheks(Where{})
}
func (r *GqlResolver) DocumentVozvratMaterialovIzProizvodstva(ctx context.Context, args DocumentVozvratMaterialovIzProizvodstvaArgs) (*DocumentVozvratMaterialovIzProizvodstva, error) {
	return r.Client.DocumentVozvratMaterialovIzProizvodstva(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratMaterialovIzProizvodstvas(ctx context.Context, args DocumentVozvratMaterialovIzProizvodstvasArgs) (*[]DocumentVozvratMaterialovIzProizvodstva, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratMaterialovIzProizvodstvas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratMaterialovIzProizvodstvas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratMaterialovIzProizvodstvas(Where{})
}
func (r *GqlResolver) DocumentVozvratMaterialovIzProizvodstvaTovary(ctx context.Context, args DocumentVozvratMaterialovIzProizvodstvaTovaryArgs) (*DocumentVozvratMaterialovIzProizvodstvaTovary, error) {
	return r.Client.DocumentVozvratMaterialovIzProizvodstvaTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratMaterialovIzProizvodstvaTovarys(ctx context.Context, args DocumentVozvratMaterialovIzProizvodstvaTovarysArgs) (*[]DocumentVozvratMaterialovIzProizvodstvaTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratMaterialovIzProizvodstvaTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratMaterialovIzProizvodstvaTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratMaterialovIzProizvodstvaTovarys(Where{})
}
func (r *GqlResolver) DocumentPereotsenkaTovarovOtdannykhNaKomissiiu(ctx context.Context, args DocumentPereotsenkaTovarovOtdannykhNaKomissiiuArgs) (*DocumentPereotsenkaTovarovOtdannykhNaKomissiiu, error) {
	return r.Client.DocumentPereotsenkaTovarovOtdannykhNaKomissiiu(args.Key, nil)
}
func (r *GqlResolver) DocumentPereotsenkaTovarovOtdannykhNaKomissiius(ctx context.Context, args DocumentPereotsenkaTovarovOtdannykhNaKomissiiusArgs) (*[]DocumentPereotsenkaTovarovOtdannykhNaKomissiiu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPereotsenkaTovarovOtdannykhNaKomissiius(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPereotsenkaTovarovOtdannykhNaKomissiius(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPereotsenkaTovarovOtdannykhNaKomissiius(Where{})
}
func (r *GqlResolver) DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary(ctx context.Context, args DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovaryArgs) (*DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary, error) {
	return r.Client.DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovarys(ctx context.Context, args DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovarysArgs) (*[]DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPereotsenkaTovarovOtdannykhNaKomissiiuTovarys(Where{})
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSN(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNArgs) (*DocumentVvodNachalnykhOstatkovPoRaskhodamUSN, error) {
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSN(args.Key, nil)
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSNs(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNsArgs) (*[]DocumentVvodNachalnykhOstatkovPoRaskhodamUSN, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNs(Where{})
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamiArgs) (*DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami, error) {
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami(args.Key, nil)
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamis(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamisArgs) (*[]DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPokupateliamis(Where{})
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyeArgs) (*DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye, error) {
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye(args.Key, nil)
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyes(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyesArgs) (*[]DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannye, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyes(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryProdannyes(Where{})
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamiArgs) (*DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami, error) {
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami(args.Key, nil)
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamis(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamisArgs) (*[]DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikami, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNVzaimoraschetySPostavshchikamis(Where{})
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhArgs) (*DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh, error) {
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh(args.Key, nil)
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhs(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhsArgs) (*[]DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVvodNachalnykhOstatkovPoRaskhodamUSNTovaryNaOstatkakhs(Where{})
}
func (r *GqlResolver) DocumentGTDImport(ctx context.Context, args DocumentGTDImportArgs) (*DocumentGTDImport, error) {
	return r.Client.DocumentGTDImport(args.Key, nil)
}
func (r *GqlResolver) DocumentGTDImports(ctx context.Context, args DocumentGTDImportsArgs) (*[]DocumentGTDImport, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentGTDImports(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentGTDImports(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentGTDImports(Where{})
}
func (r *GqlResolver) DocumentGTDImportRazdely(ctx context.Context, args DocumentGTDImportRazdelyArgs) (*DocumentGTDImportRazdely, error) {
	return r.Client.DocumentGTDImportRazdely(args.Key, nil)
}
func (r *GqlResolver) DocumentGTDImportRazdelys(ctx context.Context, args DocumentGTDImportRazdelysArgs) (*[]DocumentGTDImportRazdely, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentGTDImportRazdelys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentGTDImportRazdelys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentGTDImportRazdelys(Where{})
}
func (r *GqlResolver) DocumentGTDImportTovary(ctx context.Context, args DocumentGTDImportTovaryArgs) (*DocumentGTDImportTovary, error) {
	return r.Client.DocumentGTDImportTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentGTDImportTovarys(ctx context.Context, args DocumentGTDImportTovarysArgs) (*[]DocumentGTDImportTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentGTDImportTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentGTDImportTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentGTDImportTovarys(Where{})
}
func (r *GqlResolver) DocumentAktSverki(ctx context.Context, args DocumentAktSverkiArgs) (*DocumentAktSverki, error) {
	return r.Client.DocumentAktSverki(args.Key, nil)
}
func (r *GqlResolver) DocumentAktSverkis(ctx context.Context, args DocumentAktSverkisArgs) (*[]DocumentAktSverki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAktSverkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAktSverkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAktSverkis(Where{})
}
func (r *GqlResolver) DocumentAktSverkiTovary(ctx context.Context, args DocumentAktSverkiTovaryArgs) (*DocumentAktSverkiTovary, error) {
	return r.Client.DocumentAktSverkiTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentAktSverkiTovarys(ctx context.Context, args DocumentAktSverkiTovarysArgs) (*[]DocumentAktSverkiTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAktSverkiTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAktSverkiTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAktSverkiTovarys(Where{})
}
func (r *GqlResolver) CatalogFaily(ctx context.Context, args CatalogFailyArgs) (*CatalogFaily, error) {
	return r.Client.CatalogFaily(args.Key, nil)
}
func (r *GqlResolver) CatalogFailys(ctx context.Context, args CatalogFailysArgs) (*[]CatalogFaily, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFailys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFailys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFailys(Where{})
}
func (r *GqlResolver) CatalogFailyDopolnitelnyeRekvizity(ctx context.Context, args CatalogFailyDopolnitelnyeRekvizityArgs) (*CatalogFailyDopolnitelnyeRekvizity, error) {
	return r.Client.CatalogFailyDopolnitelnyeRekvizity(args.Key, nil)
}
func (r *GqlResolver) CatalogFailyDopolnitelnyeRekvizitys(ctx context.Context, args CatalogFailyDopolnitelnyeRekvizitysArgs) (*[]CatalogFailyDopolnitelnyeRekvizity, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFailyDopolnitelnyeRekvizitys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFailyDopolnitelnyeRekvizitys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFailyDopolnitelnyeRekvizitys(Where{})
}
func (r *GqlResolver) CatalogFailySertifikatyShifrovaniia(ctx context.Context, args CatalogFailySertifikatyShifrovaniiaArgs) (*CatalogFailySertifikatyShifrovaniia, error) {
	return r.Client.CatalogFailySertifikatyShifrovaniia(args.Key, nil)
}
func (r *GqlResolver) CatalogFailySertifikatyShifrovaniias(ctx context.Context, args CatalogFailySertifikatyShifrovaniiasArgs) (*[]CatalogFailySertifikatyShifrovaniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFailySertifikatyShifrovaniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFailySertifikatyShifrovaniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFailySertifikatyShifrovaniias(Where{})
}
func (r *GqlResolver) CatalogUchetnyeZapisiElektronnoiPochty(ctx context.Context, args CatalogUchetnyeZapisiElektronnoiPochtyArgs) (*CatalogUchetnyeZapisiElektronnoiPochty, error) {
	return r.Client.CatalogUchetnyeZapisiElektronnoiPochty(args.Key, nil)
}
func (r *GqlResolver) CatalogUchetnyeZapisiElektronnoiPochtys(ctx context.Context, args CatalogUchetnyeZapisiElektronnoiPochtysArgs) (*[]CatalogUchetnyeZapisiElektronnoiPochty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUchetnyeZapisiElektronnoiPochtys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUchetnyeZapisiElektronnoiPochtys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUchetnyeZapisiElektronnoiPochtys(Where{})
}
func (r *GqlResolver) CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi(ctx context.Context, args CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisiArgs) (*CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi, error) {
	return r.Client.CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi(args.Key, nil)
}
func (r *GqlResolver) CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisis(ctx context.Context, args CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisisArgs) (*[]CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUchetnyeZapisiElektronnoiPochtyDostupKUchetnoiZapisis(Where{})
}
func (r *GqlResolver) DocumentPlaniruemoePostuplenieDenezhnykhSredstv(ctx context.Context, args DocumentPlaniruemoePostuplenieDenezhnykhSredstvArgs) (*DocumentPlaniruemoePostuplenieDenezhnykhSredstv, error) {
	return r.Client.DocumentPlaniruemoePostuplenieDenezhnykhSredstv(args.Key, nil)
}
func (r *GqlResolver) DocumentPlaniruemoePostuplenieDenezhnykhSredstvs(ctx context.Context, args DocumentPlaniruemoePostuplenieDenezhnykhSredstvsArgs) (*[]DocumentPlaniruemoePostuplenieDenezhnykhSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlaniruemoePostuplenieDenezhnykhSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlaniruemoePostuplenieDenezhnykhSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlaniruemoePostuplenieDenezhnykhSredstvs(Where{})
}
func (r *GqlResolver) DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha(ctx context.Context, args DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhaArgs) (*DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha, error) {
	return r.Client.DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(ctx context.Context, args DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhasArgs) (*[]DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlaniruemoePostuplenieDenezhnykhSredstvRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentPreiskurantTsenNaKamni(ctx context.Context, args DocumentPreiskurantTsenNaKamniArgs) (*DocumentPreiskurantTsenNaKamni, error) {
	return r.Client.DocumentPreiskurantTsenNaKamni(args.Key, nil)
}
func (r *GqlResolver) DocumentPreiskurantTsenNaKamnis(ctx context.Context, args DocumentPreiskurantTsenNaKamnisArgs) (*[]DocumentPreiskurantTsenNaKamni, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPreiskurantTsenNaKamnis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPreiskurantTsenNaKamnis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPreiskurantTsenNaKamnis(Where{})
}
func (r *GqlResolver) Purchase(ctx context.Context, args PurchaseArgs) (*Purchase, error) {
	return r.Client.Purchase(args.Key, nil)
}
func (r *GqlResolver) Purchases(ctx context.Context, args PurchasesArgs) (*[]Purchase, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Purchases(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Purchases(Where{Filter: *args.Filter})
	}
	return r.Client.Purchases(Where{})
}
func (r *GqlResolver) DocumentSkupkaTovarovTovary(ctx context.Context, args DocumentSkupkaTovarovTovaryArgs) (*DocumentSkupkaTovarovTovary, error) {
	return r.Client.DocumentSkupkaTovarovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentSkupkaTovarovTovarys(ctx context.Context, args DocumentSkupkaTovarovTovarysArgs) (*[]DocumentSkupkaTovarovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSkupkaTovarovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSkupkaTovarovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSkupkaTovarovTovarys(Where{})
}
func (r *GqlResolver) DocumentSchetFakturaPoluchennyi(ctx context.Context, args DocumentSchetFakturaPoluchennyiArgs) (*DocumentSchetFakturaPoluchennyi, error) {
	return r.Client.DocumentSchetFakturaPoluchennyi(args.Key, nil)
}
func (r *GqlResolver) DocumentSchetFakturaPoluchennyis(ctx context.Context, args DocumentSchetFakturaPoluchennyisArgs) (*[]DocumentSchetFakturaPoluchennyi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSchetFakturaPoluchennyis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSchetFakturaPoluchennyis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSchetFakturaPoluchennyis(Where{})
}
func (r *GqlResolver) DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam(ctx context.Context, args DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliamArgs) (*DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam, error) {
	return r.Client.DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam(args.Key, nil)
}
func (r *GqlResolver) DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliams(ctx context.Context, args DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliamsArgs) (*[]DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliams(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSchetFakturaPoluchennyiSchetaFakturyVydannyePokupateliams(Where{})
}
func (r *GqlResolver) DocumentAktKhimicheskogoAnalizaMetalla(ctx context.Context, args DocumentAktKhimicheskogoAnalizaMetallaArgs) (*DocumentAktKhimicheskogoAnalizaMetalla, error) {
	return r.Client.DocumentAktKhimicheskogoAnalizaMetalla(args.Key, nil)
}
func (r *GqlResolver) DocumentAktKhimicheskogoAnalizaMetallas(ctx context.Context, args DocumentAktKhimicheskogoAnalizaMetallasArgs) (*[]DocumentAktKhimicheskogoAnalizaMetalla, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAktKhimicheskogoAnalizaMetallas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAktKhimicheskogoAnalizaMetallas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAktKhimicheskogoAnalizaMetallas(Where{})
}
func (r *GqlResolver) CatalogfmKartochkaKontragenta(ctx context.Context, args CatalogfmKartochkaKontragentaArgs) (*CatalogfmKartochkaKontragenta, error) {
	return r.Client.CatalogfmKartochkaKontragenta(args.Key, nil)
}
func (r *GqlResolver) CatalogfmKartochkaKontragentas(ctx context.Context, args CatalogfmKartochkaKontragentasArgs) (*[]CatalogfmKartochkaKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogfmKartochkaKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogfmKartochkaKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogfmKartochkaKontragentas(Where{})
}
func (r *GqlResolver) CatalogfmKartochkaKontragentaDannyeKontragenta(ctx context.Context, args CatalogfmKartochkaKontragentaDannyeKontragentaArgs) (*CatalogfmKartochkaKontragentaDannyeKontragenta, error) {
	return r.Client.CatalogfmKartochkaKontragentaDannyeKontragenta(args.Key, nil)
}
func (r *GqlResolver) CatalogfmKartochkaKontragentaDannyeKontragentas(ctx context.Context, args CatalogfmKartochkaKontragentaDannyeKontragentasArgs) (*[]CatalogfmKartochkaKontragentaDannyeKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogfmKartochkaKontragentaDannyeKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogfmKartochkaKontragentaDannyeKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogfmKartochkaKontragentaDannyeKontragentas(Where{})
}
func (r *GqlResolver) DocumentSpisanieProsrochennykhSertifikatov(ctx context.Context, args DocumentSpisanieProsrochennykhSertifikatovArgs) (*DocumentSpisanieProsrochennykhSertifikatov, error) {
	return r.Client.DocumentSpisanieProsrochennykhSertifikatov(args.Key, nil)
}
func (r *GqlResolver) DocumentSpisanieProsrochennykhSertifikatovs(ctx context.Context, args DocumentSpisanieProsrochennykhSertifikatovsArgs) (*[]DocumentSpisanieProsrochennykhSertifikatov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSpisanieProsrochennykhSertifikatovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSpisanieProsrochennykhSertifikatovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSpisanieProsrochennykhSertifikatovs(Where{})
}
func (r *GqlResolver) DocumentSpisanieProsrochennykhSertifikatovSertifikaty(ctx context.Context, args DocumentSpisanieProsrochennykhSertifikatovSertifikatyArgs) (*DocumentSpisanieProsrochennykhSertifikatovSertifikaty, error) {
	return r.Client.DocumentSpisanieProsrochennykhSertifikatovSertifikaty(args.Key, nil)
}
func (r *GqlResolver) DocumentSpisanieProsrochennykhSertifikatovSertifikatys(ctx context.Context, args DocumentSpisanieProsrochennykhSertifikatovSertifikatysArgs) (*[]DocumentSpisanieProsrochennykhSertifikatovSertifikaty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSpisanieProsrochennykhSertifikatovSertifikatys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSpisanieProsrochennykhSertifikatovSertifikatys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSpisanieProsrochennykhSertifikatovSertifikatys(Where{})
}
func (r *GqlResolver) DocumentZakrytieAvansovPoGrafikuPlatezhei(ctx context.Context, args DocumentZakrytieAvansovPoGrafikuPlatezheiArgs) (*DocumentZakrytieAvansovPoGrafikuPlatezhei, error) {
	return r.Client.DocumentZakrytieAvansovPoGrafikuPlatezhei(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytieAvansovPoGrafikuPlatezheis(ctx context.Context, args DocumentZakrytieAvansovPoGrafikuPlatezheisArgs) (*[]DocumentZakrytieAvansovPoGrafikuPlatezhei, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytieAvansovPoGrafikuPlatezheis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytieAvansovPoGrafikuPlatezheis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytieAvansovPoGrafikuPlatezheis(Where{})
}
func (r *GqlResolver) DocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty(ctx context.Context, args DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentyArgs) (*DocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty, error) {
	return r.Client.DocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentys(ctx context.Context, args DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentysArgs) (*[]DocumentZakrytieAvansovPoGrafikuPlatezheiKontragenty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytieAvansovPoGrafikuPlatezheiKontragentys(Where{})
}
func (r *GqlResolver) CatalogTipySistemNalogooblozheniiaKKT(ctx context.Context, args CatalogTipySistemNalogooblozheniiaKKTArgs) (*CatalogTipySistemNalogooblozheniiaKKT, error) {
	return r.Client.CatalogTipySistemNalogooblozheniiaKKT(args.Key, nil)
}
func (r *GqlResolver) CatalogTipySistemNalogooblozheniiaKKTs(ctx context.Context, args CatalogTipySistemNalogooblozheniiaKKTsArgs) (*[]CatalogTipySistemNalogooblozheniiaKKT, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTipySistemNalogooblozheniiaKKTs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTipySistemNalogooblozheniiaKKTs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTipySistemNalogooblozheniiaKKTs(Where{})
}
func (r *GqlResolver) DocumentAkkreditivPeredannyi(ctx context.Context, args DocumentAkkreditivPeredannyiArgs) (*DocumentAkkreditivPeredannyi, error) {
	return r.Client.DocumentAkkreditivPeredannyi(args.Key, nil)
}
func (r *GqlResolver) DocumentAkkreditivPeredannyis(ctx context.Context, args DocumentAkkreditivPeredannyisArgs) (*[]DocumentAkkreditivPeredannyi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAkkreditivPeredannyis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAkkreditivPeredannyis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAkkreditivPeredannyis(Where{})
}
func (r *GqlResolver) DocumentAkkreditivPeredannyiRasshifrovkaPlatezha(ctx context.Context, args DocumentAkkreditivPeredannyiRasshifrovkaPlatezhaArgs) (*DocumentAkkreditivPeredannyiRasshifrovkaPlatezha, error) {
	return r.Client.DocumentAkkreditivPeredannyiRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentAkkreditivPeredannyiRasshifrovkaPlatezhas(ctx context.Context, args DocumentAkkreditivPeredannyiRasshifrovkaPlatezhasArgs) (*[]DocumentAkkreditivPeredannyiRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAkkreditivPeredannyiRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAkkreditivPeredannyiRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAkkreditivPeredannyiRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentAkkreditivPeredannyiRekvizityKontragenta(ctx context.Context, args DocumentAkkreditivPeredannyiRekvizityKontragentaArgs) (*DocumentAkkreditivPeredannyiRekvizityKontragenta, error) {
	return r.Client.DocumentAkkreditivPeredannyiRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentAkkreditivPeredannyiRekvizityKontragentas(ctx context.Context, args DocumentAkkreditivPeredannyiRekvizityKontragentasArgs) (*[]DocumentAkkreditivPeredannyiRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAkkreditivPeredannyiRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAkkreditivPeredannyiRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAkkreditivPeredannyiRekvizityKontragentas(Where{})
}
func (r *GqlResolver) Supplier(ctx context.Context, args SupplierArgs) (*Supplier, error) {
	return r.Client.Supplier(args.Key, nil)
}
func (r *GqlResolver) Suppliers(ctx context.Context, args SuppliersArgs) (*[]Supplier, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Suppliers(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Suppliers(Where{Filter: *args.Filter})
	}
	return r.Client.Suppliers(Where{})
}
func (r *GqlResolver) CatalogKontragentyVidyDeiatelnosti(ctx context.Context, args CatalogKontragentyVidyDeiatelnostiArgs) (*CatalogKontragentyVidyDeiatelnosti, error) {
	return r.Client.CatalogKontragentyVidyDeiatelnosti(args.Key, nil)
}
func (r *GqlResolver) CatalogKontragentyVidyDeiatelnostis(ctx context.Context, args CatalogKontragentyVidyDeiatelnostisArgs) (*[]CatalogKontragentyVidyDeiatelnosti, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKontragentyVidyDeiatelnostis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKontragentyVidyDeiatelnostis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKontragentyVidyDeiatelnostis(Where{})
}
func (r *GqlResolver) DocumentInformatsionnoeSoobshchenie(ctx context.Context, args DocumentInformatsionnoeSoobshchenieArgs) (*DocumentInformatsionnoeSoobshchenie, error) {
	return r.Client.DocumentInformatsionnoeSoobshchenie(args.Key, nil)
}
func (r *GqlResolver) DocumentInformatsionnoeSoobshchenies(ctx context.Context, args DocumentInformatsionnoeSoobshcheniesArgs) (*[]DocumentInformatsionnoeSoobshchenie, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInformatsionnoeSoobshchenies(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInformatsionnoeSoobshchenies(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInformatsionnoeSoobshchenies(Where{})
}
func (r *GqlResolver) CatalogVlozheniiaElektronnykhPisem(ctx context.Context, args CatalogVlozheniiaElektronnykhPisemArgs) (*CatalogVlozheniiaElektronnykhPisem, error) {
	return r.Client.CatalogVlozheniiaElektronnykhPisem(args.Key, nil)
}
func (r *GqlResolver) CatalogVlozheniiaElektronnykhPisems(ctx context.Context, args CatalogVlozheniiaElektronnykhPisemsArgs) (*[]CatalogVlozheniiaElektronnykhPisem, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVlozheniiaElektronnykhPisems(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVlozheniiaElektronnykhPisems(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVlozheniiaElektronnykhPisems(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovanieVystavlennoe(ctx context.Context, args DocumentPlatezhnoeTrebovanieVystavlennoeArgs) (*DocumentPlatezhnoeTrebovanieVystavlennoe, error) {
	return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoe(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovanieVystavlennoes(ctx context.Context, args DocumentPlatezhnoeTrebovanieVystavlennoesArgs) (*[]DocumentPlatezhnoeTrebovanieVystavlennoe, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoes(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoes(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha(ctx context.Context, args DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhaArgs) (*DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha, error) {
	return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhas(ctx context.Context, args DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhasArgs) (*[]DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoeRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta(ctx context.Context, args DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentaArgs) (*DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta, error) {
	return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentas(ctx context.Context, args DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentasArgs) (*[]DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoeTrebovanieVystavlennoeRekvizityKontragentas(Where{})
}
func (r *GqlResolver) DocumentMarketingovaiaAktsiia(ctx context.Context, args DocumentMarketingovaiaAktsiiaArgs) (*DocumentMarketingovaiaAktsiia, error) {
	return r.Client.DocumentMarketingovaiaAktsiia(args.Key, nil)
}
func (r *GqlResolver) DocumentMarketingovaiaAktsiias(ctx context.Context, args DocumentMarketingovaiaAktsiiasArgs) (*[]DocumentMarketingovaiaAktsiia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentMarketingovaiaAktsiias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentMarketingovaiaAktsiias(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentMarketingovaiaAktsiias(Where{})
}
func (r *GqlResolver) DocumentMarketingovaiaAktsiiaMagaziny(ctx context.Context, args DocumentMarketingovaiaAktsiiaMagazinyArgs) (*DocumentMarketingovaiaAktsiiaMagaziny, error) {
	return r.Client.DocumentMarketingovaiaAktsiiaMagaziny(args.Key, nil)
}
func (r *GqlResolver) DocumentMarketingovaiaAktsiiaMagazinys(ctx context.Context, args DocumentMarketingovaiaAktsiiaMagazinysArgs) (*[]DocumentMarketingovaiaAktsiiaMagaziny, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentMarketingovaiaAktsiiaMagazinys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentMarketingovaiaAktsiiaMagazinys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentMarketingovaiaAktsiiaMagazinys(Where{})
}
func (r *GqlResolver) DocumentMarketingovaiaAktsiiaSkidkiNatsenki(ctx context.Context, args DocumentMarketingovaiaAktsiiaSkidkiNatsenkiArgs) (*DocumentMarketingovaiaAktsiiaSkidkiNatsenki, error) {
	return r.Client.DocumentMarketingovaiaAktsiiaSkidkiNatsenki(args.Key, nil)
}
func (r *GqlResolver) DocumentMarketingovaiaAktsiiaSkidkiNatsenkis(ctx context.Context, args DocumentMarketingovaiaAktsiiaSkidkiNatsenkisArgs) (*[]DocumentMarketingovaiaAktsiiaSkidkiNatsenki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentMarketingovaiaAktsiiaSkidkiNatsenkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentMarketingovaiaAktsiiaSkidkiNatsenkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentMarketingovaiaAktsiiaSkidkiNatsenkis(Where{})
}
func (r *GqlResolver) DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa(ctx context.Context, args DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupaArgs) (*DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa, error) {
	return r.Client.DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa(args.Key, nil)
}
func (r *GqlResolver) DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupas(ctx context.Context, args DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupasArgs) (*[]DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentMarketingovaiaAktsiiaNaboryZnacheniiDostupas(Where{})
}
func (r *GqlResolver) CatalogStsenariiObmenovDannymi(ctx context.Context, args CatalogStsenariiObmenovDannymiArgs) (*CatalogStsenariiObmenovDannymi, error) {
	return r.Client.CatalogStsenariiObmenovDannymi(args.Key, nil)
}
func (r *GqlResolver) CatalogStsenariiObmenovDannymis(ctx context.Context, args CatalogStsenariiObmenovDannymisArgs) (*[]CatalogStsenariiObmenovDannymi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogStsenariiObmenovDannymis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogStsenariiObmenovDannymis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogStsenariiObmenovDannymis(Where{})
}
func (r *GqlResolver) CatalogStsenariiObmenovDannymiNastroikiObmena(ctx context.Context, args CatalogStsenariiObmenovDannymiNastroikiObmenaArgs) (*CatalogStsenariiObmenovDannymiNastroikiObmena, error) {
	return r.Client.CatalogStsenariiObmenovDannymiNastroikiObmena(args.Key, nil)
}
func (r *GqlResolver) CatalogStsenariiObmenovDannymiNastroikiObmenas(ctx context.Context, args CatalogStsenariiObmenovDannymiNastroikiObmenasArgs) (*[]CatalogStsenariiObmenovDannymiNastroikiObmena, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogStsenariiObmenovDannymiNastroikiObmenas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogStsenariiObmenovDannymiNastroikiObmenas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogStsenariiObmenovDannymiNastroikiObmenas(Where{})
}
func (r *GqlResolver) Item(ctx context.Context, args ItemArgs) (*Item, error) {
	return r.Client.Item(args.Key, nil)
}
func (r *GqlResolver) Items(ctx context.Context, args ItemsArgs) (*[]Item, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Items(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Items(Where{Filter: *args.Filter})
	}
	return r.Client.Items(Where{})
}
func (r *GqlResolver) CatalogNomenklaturaSostavLigatury(ctx context.Context, args CatalogNomenklaturaSostavLigaturyArgs) (*CatalogNomenklaturaSostavLigatury, error) {
	return r.Client.CatalogNomenklaturaSostavLigatury(args.Key, nil)
}
func (r *GqlResolver) CatalogNomenklaturaSostavLigaturys(ctx context.Context, args CatalogNomenklaturaSostavLigaturysArgs) (*[]CatalogNomenklaturaSostavLigatury, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNomenklaturaSostavLigaturys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNomenklaturaSostavLigaturys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNomenklaturaSostavLigaturys(Where{})
}
func (r *GqlResolver) DocumentOpros(ctx context.Context, args DocumentOprosArgs) (*DocumentOpros, error) {
	return r.Client.DocumentOpros(args.Key, nil)
}
func (r *GqlResolver) DocumentOpross(ctx context.Context, args DocumentOprossArgs) (*[]DocumentOpros, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOpross(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOpross(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOpross(Where{})
}
func (r *GqlResolver) DocumentOprosVoprosy(ctx context.Context, args DocumentOprosVoprosyArgs) (*DocumentOprosVoprosy, error) {
	return r.Client.DocumentOprosVoprosy(args.Key, nil)
}
func (r *GqlResolver) DocumentOprosVoprosys(ctx context.Context, args DocumentOprosVoprosysArgs) (*[]DocumentOprosVoprosy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOprosVoprosys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOprosVoprosys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOprosVoprosys(Where{})
}
func (r *GqlResolver) DocumentOprosSostavnoiOtvet(ctx context.Context, args DocumentOprosSostavnoiOtvetArgs) (*DocumentOprosSostavnoiOtvet, error) {
	return r.Client.DocumentOprosSostavnoiOtvet(args.Key, nil)
}
func (r *GqlResolver) DocumentOprosSostavnoiOtvets(ctx context.Context, args DocumentOprosSostavnoiOtvetsArgs) (*[]DocumentOprosSostavnoiOtvet, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOprosSostavnoiOtvets(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOprosSostavnoiOtvets(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOprosSostavnoiOtvets(Where{})
}
func (r *GqlResolver) CatalogGruppyPoluchateleiSkidki(ctx context.Context, args CatalogGruppyPoluchateleiSkidkiArgs) (*CatalogGruppyPoluchateleiSkidki, error) {
	return r.Client.CatalogGruppyPoluchateleiSkidki(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyPoluchateleiSkidkis(ctx context.Context, args CatalogGruppyPoluchateleiSkidkisArgs) (*[]CatalogGruppyPoluchateleiSkidki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyPoluchateleiSkidkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyPoluchateleiSkidkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyPoluchateleiSkidkis(Where{})
}
func (r *GqlResolver) Reassessment(ctx context.Context, args ReassessmentArgs) (*Reassessment, error) {
	return r.Client.Reassessment(args.Key, nil)
}
func (r *GqlResolver) Reassessments(ctx context.Context, args ReassessmentsArgs) (*[]Reassessment, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Reassessments(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Reassessments(Where{Filter: *args.Filter})
	}
	return r.Client.Reassessments(Where{})
}
func (r *GqlResolver) DocumentPereotsenkaTovarovVNTTTovary(ctx context.Context, args DocumentPereotsenkaTovarovVNTTTovaryArgs) (*DocumentPereotsenkaTovarovVNTTTovary, error) {
	return r.Client.DocumentPereotsenkaTovarovVNTTTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPereotsenkaTovarovVNTTTovarys(ctx context.Context, args DocumentPereotsenkaTovarovVNTTTovarysArgs) (*[]DocumentPereotsenkaTovarovVNTTTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPereotsenkaTovarovVNTTTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPereotsenkaTovarovVNTTTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPereotsenkaTovarovVNTTTovarys(Where{})
}
func (r *GqlResolver) CatalogTomaKhraneniiaFailov(ctx context.Context, args CatalogTomaKhraneniiaFailovArgs) (*CatalogTomaKhraneniiaFailov, error) {
	return r.Client.CatalogTomaKhraneniiaFailov(args.Key, nil)
}
func (r *GqlResolver) CatalogTomaKhraneniiaFailovs(ctx context.Context, args CatalogTomaKhraneniiaFailovsArgs) (*[]CatalogTomaKhraneniiaFailov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTomaKhraneniiaFailovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTomaKhraneniiaFailovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTomaKhraneniiaFailovs(Where{})
}
func (r *GqlResolver) DocumentJournalProizvodstvennyeDokumenty(ctx context.Context, args DocumentJournalProizvodstvennyeDokumentyArgs) (*DocumentJournalProizvodstvennyeDokumenty, error) {
	return r.Client.DocumentJournalProizvodstvennyeDokumenty(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalProizvodstvennyeDokumentys(ctx context.Context, args DocumentJournalProizvodstvennyeDokumentysArgs) (*[]DocumentJournalProizvodstvennyeDokumenty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalProizvodstvennyeDokumentys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalProizvodstvennyeDokumentys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalProizvodstvennyeDokumentys(Where{})
}
func (r *GqlResolver) DocumentIzmeneniePravDostupa(ctx context.Context, args DocumentIzmeneniePravDostupaArgs) (*DocumentIzmeneniePravDostupa, error) {
	return r.Client.DocumentIzmeneniePravDostupa(args.Key, nil)
}
func (r *GqlResolver) DocumentIzmeneniePravDostupas(ctx context.Context, args DocumentIzmeneniePravDostupasArgs) (*[]DocumentIzmeneniePravDostupa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentIzmeneniePravDostupas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentIzmeneniePravDostupas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentIzmeneniePravDostupas(Where{})
}
func (r *GqlResolver) CatalogNastroikaAssortimentnoiMatritsy(ctx context.Context, args CatalogNastroikaAssortimentnoiMatritsyArgs) (*CatalogNastroikaAssortimentnoiMatritsy, error) {
	return r.Client.CatalogNastroikaAssortimentnoiMatritsy(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikaAssortimentnoiMatritsys(ctx context.Context, args CatalogNastroikaAssortimentnoiMatritsysArgs) (*[]CatalogNastroikaAssortimentnoiMatritsy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikaAssortimentnoiMatritsys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikaAssortimentnoiMatritsys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikaAssortimentnoiMatritsys(Where{})
}
func (r *GqlResolver) CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp(ctx context.Context, args CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGruppArgs) (*CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp, error) {
	return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupps(ctx context.Context, args CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGruppsArgs) (*[]CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupp, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupps(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupps(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhGrupps(Where{})
}
func (r *GqlResolver) CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii(ctx context.Context, args CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriiArgs) (*CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii, error) {
	return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriis(ctx context.Context, args CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriisArgs) (*[]CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategorii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhKategoriis(Where{})
}
func (r *GqlResolver) CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii(ctx context.Context, args CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiiArgs) (*CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii, error) {
	return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiis(ctx context.Context, args CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiisArgs) (*[]CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikaAssortimentnoiMatritsyNastroikaTovarnykhPozitsiis(Where{})
}
func (r *GqlResolver) DocumentJournalDokumentyKontragentov(ctx context.Context, args DocumentJournalDokumentyKontragentovArgs) (*DocumentJournalDokumentyKontragentov, error) {
	return r.Client.DocumentJournalDokumentyKontragentov(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalDokumentyKontragentovs(ctx context.Context, args DocumentJournalDokumentyKontragentovsArgs) (*[]DocumentJournalDokumentyKontragentov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalDokumentyKontragentovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalDokumentyKontragentovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalDokumentyKontragentovs(Where{})
}
func (r *GqlResolver) MoveInstance(ctx context.Context, args MoveInstanceArgs) (*MoveInstance, error) {
	return r.Client.MoveInstance(args.Key, nil)
}
func (r *GqlResolver) MoveInstances(ctx context.Context, args MoveInstancesArgs) (*[]MoveInstance, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.MoveInstances(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.MoveInstances(Where{Filter: *args.Filter})
	}
	return r.Client.MoveInstances(Where{})
}
func (r *GqlResolver) DocumentPeremeshchenieTovarovSertifikaty(ctx context.Context, args DocumentPeremeshchenieTovarovSertifikatyArgs) (*DocumentPeremeshchenieTovarovSertifikaty, error) {
	return r.Client.DocumentPeremeshchenieTovarovSertifikaty(args.Key, nil)
}
func (r *GqlResolver) DocumentPeremeshchenieTovarovSertifikatys(ctx context.Context, args DocumentPeremeshchenieTovarovSertifikatysArgs) (*[]DocumentPeremeshchenieTovarovSertifikaty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPeremeshchenieTovarovSertifikatys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPeremeshchenieTovarovSertifikatys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPeremeshchenieTovarovSertifikatys(Where{})
}
func (r *GqlResolver) DocumentPeremeshchenieTovarovTovary(ctx context.Context, args DocumentPeremeshchenieTovarovTovaryArgs) (*DocumentPeremeshchenieTovarovTovary, error) {
	return r.Client.DocumentPeremeshchenieTovarovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPeremeshchenieTovarovTovarys(ctx context.Context, args DocumentPeremeshchenieTovarovTovarysArgs) (*[]DocumentPeremeshchenieTovarovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPeremeshchenieTovarovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPeremeshchenieTovarovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPeremeshchenieTovarovTovarys(Where{})
}
func (r *GqlResolver) DocumentPeremeshchenieTovarovSpisokZaiavok(ctx context.Context, args DocumentPeremeshchenieTovarovSpisokZaiavokArgs) (*DocumentPeremeshchenieTovarovSpisokZaiavok, error) {
	return r.Client.DocumentPeremeshchenieTovarovSpisokZaiavok(args.Key, nil)
}
func (r *GqlResolver) DocumentPeremeshchenieTovarovSpisokZaiavoks(ctx context.Context, args DocumentPeremeshchenieTovarovSpisokZaiavoksArgs) (*[]DocumentPeremeshchenieTovarovSpisokZaiavok, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPeremeshchenieTovarovSpisokZaiavoks(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPeremeshchenieTovarovSpisokZaiavoks(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPeremeshchenieTovarovSpisokZaiavoks(Where{})
}
func (r *GqlResolver) DocumentZakrytieZaiavokNaRaskhodovanieSredstv(ctx context.Context, args DocumentZakrytieZaiavokNaRaskhodovanieSredstvArgs) (*DocumentZakrytieZaiavokNaRaskhodovanieSredstv, error) {
	return r.Client.DocumentZakrytieZaiavokNaRaskhodovanieSredstv(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytieZaiavokNaRaskhodovanieSredstvs(ctx context.Context, args DocumentZakrytieZaiavokNaRaskhodovanieSredstvsArgs) (*[]DocumentZakrytieZaiavokNaRaskhodovanieSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytieZaiavokNaRaskhodovanieSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytieZaiavokNaRaskhodovanieSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytieZaiavokNaRaskhodovanieSredstvs(Where{})
}
func (r *GqlResolver) DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv(ctx context.Context, args DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvArgs) (*DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv, error) {
	return r.Client.DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvs(ctx context.Context, args DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvsArgs) (*[]DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytieZaiavokNaRaskhodovanieSredstvZaiavkiNaRaskhodovanieSredstvs(Where{})
}
func (r *GqlResolver) MemberCard(ctx context.Context, args MemberCardArgs) (*MemberCard, error) {
	return r.Client.MemberCard(args.Key, nil)
}
func (r *GqlResolver) MemberCards(ctx context.Context, args MemberCardsArgs) (*[]MemberCard, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.MemberCards(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.MemberCards(Where{Filter: *args.Filter})
	}
	return r.Client.MemberCards(Where{})
}
func (r *GqlResolver) DocumentABCKlassifikatsiiaPokupatelei(ctx context.Context, args DocumentABCKlassifikatsiiaPokupateleiArgs) (*DocumentABCKlassifikatsiiaPokupatelei, error) {
	return r.Client.DocumentABCKlassifikatsiiaPokupatelei(args.Key, nil)
}
func (r *GqlResolver) DocumentABCKlassifikatsiiaPokupateleis(ctx context.Context, args DocumentABCKlassifikatsiiaPokupateleisArgs) (*[]DocumentABCKlassifikatsiiaPokupatelei, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentABCKlassifikatsiiaPokupateleis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentABCKlassifikatsiiaPokupateleis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentABCKlassifikatsiiaPokupateleis(Where{})
}
func (r *GqlResolver) DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov(ctx context.Context, args DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovArgs) (*DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov, error) {
	return r.Client.DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov(args.Key, nil)
}
func (r *GqlResolver) DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovs(ctx context.Context, args DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovsArgs) (*[]DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentABCKlassifikatsiiaPokupateleiTablitsaRaspredeleniiaKontragentovs(Where{})
}
func (r *GqlResolver) CatalogIdentifikatoryObieektovMetadannykh(ctx context.Context, args CatalogIdentifikatoryObieektovMetadannykhArgs) (*CatalogIdentifikatoryObieektovMetadannykh, error) {
	return r.Client.CatalogIdentifikatoryObieektovMetadannykh(args.Key, nil)
}
func (r *GqlResolver) CatalogIdentifikatoryObieektovMetadannykhs(ctx context.Context, args CatalogIdentifikatoryObieektovMetadannykhsArgs) (*[]CatalogIdentifikatoryObieektovMetadannykh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogIdentifikatoryObieektovMetadannykhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogIdentifikatoryObieektovMetadannykhs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogIdentifikatoryObieektovMetadannykhs(Where{})
}
func (r *GqlResolver) DocumentSvodnaiaInventarizatsiiaTovarovNaSklade(ctx context.Context, args DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeArgs) (*DocumentSvodnaiaInventarizatsiiaTovarovNaSklade, error) {
	return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSklade(args.Key, nil)
}
func (r *GqlResolver) DocumentSvodnaiaInventarizatsiiaTovarovNaSklades(ctx context.Context, args DocumentSvodnaiaInventarizatsiiaTovarovNaSkladesArgs) (*[]DocumentSvodnaiaInventarizatsiiaTovarovNaSklade, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSklades(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSklades(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSklades(Where{})
}
func (r *GqlResolver) DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty(ctx context.Context, args DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatyArgs) (*DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty, error) {
	return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty(args.Key, nil)
}
func (r *GqlResolver) DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatys(ctx context.Context, args DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatysArgs) (*[]DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikaty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeSertifikatys(Where{})
}
func (r *GqlResolver) DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii(ctx context.Context, args DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiiArgs) (*DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii, error) {
	return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii(args.Key, nil)
}
func (r *GqlResolver) DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(ctx context.Context, args DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiisArgs) (*[]DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSvodnaiaInventarizatsiiaTovarovNaSkladeUsloviiaProvedeniiaInventarizatsiis(Where{})
}
func (r *GqlResolver) DocumentKorrektirovkaRealizatsii(ctx context.Context, args DocumentKorrektirovkaRealizatsiiArgs) (*DocumentKorrektirovkaRealizatsii, error) {
	return r.Client.DocumentKorrektirovkaRealizatsii(args.Key, nil)
}
func (r *GqlResolver) DocumentKorrektirovkaRealizatsiis(ctx context.Context, args DocumentKorrektirovkaRealizatsiisArgs) (*[]DocumentKorrektirovkaRealizatsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKorrektirovkaRealizatsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKorrektirovkaRealizatsiis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKorrektirovkaRealizatsiis(Where{})
}
func (r *GqlResolver) DocumentKorrektirovkaRealizatsiiTovary(ctx context.Context, args DocumentKorrektirovkaRealizatsiiTovaryArgs) (*DocumentKorrektirovkaRealizatsiiTovary, error) {
	return r.Client.DocumentKorrektirovkaRealizatsiiTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentKorrektirovkaRealizatsiiTovarys(ctx context.Context, args DocumentKorrektirovkaRealizatsiiTovarysArgs) (*[]DocumentKorrektirovkaRealizatsiiTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKorrektirovkaRealizatsiiTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKorrektirovkaRealizatsiiTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKorrektirovkaRealizatsiiTovarys(Where{})
}
func (r *GqlResolver) DocumentKorrektirovkaRealizatsiiUslugi(ctx context.Context, args DocumentKorrektirovkaRealizatsiiUslugiArgs) (*DocumentKorrektirovkaRealizatsiiUslugi, error) {
	return r.Client.DocumentKorrektirovkaRealizatsiiUslugi(args.Key, nil)
}
func (r *GqlResolver) DocumentKorrektirovkaRealizatsiiUslugis(ctx context.Context, args DocumentKorrektirovkaRealizatsiiUslugisArgs) (*[]DocumentKorrektirovkaRealizatsiiUslugi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKorrektirovkaRealizatsiiUslugis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKorrektirovkaRealizatsiiUslugis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKorrektirovkaRealizatsiiUslugis(Where{})
}
func (r *GqlResolver) CatalogVidyDefektov(ctx context.Context, args CatalogVidyDefektovArgs) (*CatalogVidyDefektov, error) {
	return r.Client.CatalogVidyDefektov(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyDefektovs(ctx context.Context, args CatalogVidyDefektovsArgs) (*[]CatalogVidyDefektov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyDefektovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyDefektovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyDefektovs(Where{})
}
func (r *GqlResolver) DocumentDoverennost(ctx context.Context, args DocumentDoverennostArgs) (*DocumentDoverennost, error) {
	return r.Client.DocumentDoverennost(args.Key, nil)
}
func (r *GqlResolver) DocumentDoverennosts(ctx context.Context, args DocumentDoverennostsArgs) (*[]DocumentDoverennost, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentDoverennosts(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentDoverennosts(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentDoverennosts(Where{})
}
func (r *GqlResolver) DocumentDoverennostTovary(ctx context.Context, args DocumentDoverennostTovaryArgs) (*DocumentDoverennostTovary, error) {
	return r.Client.DocumentDoverennostTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentDoverennostTovarys(ctx context.Context, args DocumentDoverennostTovarysArgs) (*[]DocumentDoverennostTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentDoverennostTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentDoverennostTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentDoverennostTovarys(Where{})
}
func (r *GqlResolver) CatalogShablonyZapolneniiaKU(ctx context.Context, args CatalogShablonyZapolneniiaKUArgs) (*CatalogShablonyZapolneniiaKU, error) {
	return r.Client.CatalogShablonyZapolneniiaKU(args.Key, nil)
}
func (r *GqlResolver) CatalogShablonyZapolneniiaKUs(ctx context.Context, args CatalogShablonyZapolneniiaKUsArgs) (*[]CatalogShablonyZapolneniiaKU, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogShablonyZapolneniiaKUs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogShablonyZapolneniiaKUs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogShablonyZapolneniiaKUs(Where{})
}
func (r *GqlResolver) CatalogShablonyZapolneniiaKUPrazdnichnyeDni(ctx context.Context, args CatalogShablonyZapolneniiaKUPrazdnichnyeDniArgs) (*CatalogShablonyZapolneniiaKUPrazdnichnyeDni, error) {
	return r.Client.CatalogShablonyZapolneniiaKUPrazdnichnyeDni(args.Key, nil)
}
func (r *GqlResolver) CatalogShablonyZapolneniiaKUPrazdnichnyeDnis(ctx context.Context, args CatalogShablonyZapolneniiaKUPrazdnichnyeDnisArgs) (*[]CatalogShablonyZapolneniiaKUPrazdnichnyeDni, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogShablonyZapolneniiaKUPrazdnichnyeDnis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogShablonyZapolneniiaKUPrazdnichnyeDnis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogShablonyZapolneniiaKUPrazdnichnyeDnis(Where{})
}
func (r *GqlResolver) CatalogShablonyZapolneniiaKUKUNaNedeliu(ctx context.Context, args CatalogShablonyZapolneniiaKUKUNaNedeliuArgs) (*CatalogShablonyZapolneniiaKUKUNaNedeliu, error) {
	return r.Client.CatalogShablonyZapolneniiaKUKUNaNedeliu(args.Key, nil)
}
func (r *GqlResolver) CatalogShablonyZapolneniiaKUKUNaNedelius(ctx context.Context, args CatalogShablonyZapolneniiaKUKUNaNedeliusArgs) (*[]CatalogShablonyZapolneniiaKUKUNaNedeliu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogShablonyZapolneniiaKUKUNaNedelius(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogShablonyZapolneniiaKUKUNaNedelius(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogShablonyZapolneniiaKUKUNaNedelius(Where{})
}
func (r *GqlResolver) CatalogShablonyZapolneniiaKUSalony(ctx context.Context, args CatalogShablonyZapolneniiaKUSalonyArgs) (*CatalogShablonyZapolneniiaKUSalony, error) {
	return r.Client.CatalogShablonyZapolneniiaKUSalony(args.Key, nil)
}
func (r *GqlResolver) CatalogShablonyZapolneniiaKUSalonys(ctx context.Context, args CatalogShablonyZapolneniiaKUSalonysArgs) (*[]CatalogShablonyZapolneniiaKUSalony, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogShablonyZapolneniiaKUSalonys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogShablonyZapolneniiaKUSalonys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogShablonyZapolneniiaKUSalonys(Where{})
}
func (r *GqlResolver) DocumentPlanZapolneniiaVitrin(ctx context.Context, args DocumentPlanZapolneniiaVitrinArgs) (*DocumentPlanZapolneniiaVitrin, error) {
	return r.Client.DocumentPlanZapolneniiaVitrin(args.Key, nil)
}
func (r *GqlResolver) DocumentPlanZapolneniiaVitrins(ctx context.Context, args DocumentPlanZapolneniiaVitrinsArgs) (*[]DocumentPlanZapolneniiaVitrin, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlanZapolneniiaVitrins(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlanZapolneniiaVitrins(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlanZapolneniiaVitrins(Where{})
}
func (r *GqlResolver) DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin(ctx context.Context, args DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrinArgs) (*DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin, error) {
	return r.Client.DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin(args.Key, nil)
}
func (r *GqlResolver) DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrins(ctx context.Context, args DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrinsArgs) (*[]DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrin, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrins(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrins(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlanZapolneniiaVitrinPlanovoeZapolnenieVitrins(Where{})
}
func (r *GqlResolver) Instance(ctx context.Context, args InstanceArgs) (*Instance, error) {
	return r.Client.Instance(args.Key, nil)
}
func (r *GqlResolver) Instances(ctx context.Context, args InstancesArgs) (*[]Instance, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Instances(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Instances(Where{Filter: *args.Filter})
	}
	return r.Client.Instances(Where{})
}
func (r *GqlResolver) ReturnToManufacturing(ctx context.Context, args ReturnToManufacturingArgs) (*ReturnToManufacturing, error) {
	return r.Client.ReturnToManufacturing(args.Key, nil)
}
func (r *GqlResolver) ReturnToManufacturings(ctx context.Context, args ReturnToManufacturingsArgs) (*[]ReturnToManufacturing, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.ReturnToManufacturings(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.ReturnToManufacturings(Where{Filter: *args.Filter})
	}
	return r.Client.ReturnToManufacturings(Where{})
}
func (r *GqlResolver) DocumentVozvratProduktsiiVProizvodstvoTovary(ctx context.Context, args DocumentVozvratProduktsiiVProizvodstvoTovaryArgs) (*DocumentVozvratProduktsiiVProizvodstvoTovary, error) {
	return r.Client.DocumentVozvratProduktsiiVProizvodstvoTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratProduktsiiVProizvodstvoTovarys(ctx context.Context, args DocumentVozvratProduktsiiVProizvodstvoTovarysArgs) (*[]DocumentVozvratProduktsiiVProizvodstvoTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratProduktsiiVProizvodstvoTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratProduktsiiVProizvodstvoTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratProduktsiiVProizvodstvoTovarys(Where{})
}
func (r *GqlResolver) CatalogNomeraGTD(ctx context.Context, args CatalogNomeraGTDArgs) (*CatalogNomeraGTD, error) {
	return r.Client.CatalogNomeraGTD(args.Key, nil)
}
func (r *GqlResolver) CatalogNomeraGTDs(ctx context.Context, args CatalogNomeraGTDsArgs) (*[]CatalogNomeraGTD, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNomeraGTDs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNomeraGTDs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNomeraGTDs(Where{})
}
func (r *GqlResolver) CatalogNastroikiRabochegoMestaPolzovatelia(ctx context.Context, args CatalogNastroikiRabochegoMestaPolzovateliaArgs) (*CatalogNastroikiRabochegoMestaPolzovatelia, error) {
	return r.Client.CatalogNastroikiRabochegoMestaPolzovatelia(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiRabochegoMestaPolzovatelias(ctx context.Context, args CatalogNastroikiRabochegoMestaPolzovateliasArgs) (*[]CatalogNastroikiRabochegoMestaPolzovatelia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiRabochegoMestaPolzovatelias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiRabochegoMestaPolzovatelias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiRabochegoMestaPolzovatelias(Where{})
}
func (r *GqlResolver) CatalogNastroikiRabochegoMestaPolzovateliaNastroiki(ctx context.Context, args CatalogNastroikiRabochegoMestaPolzovateliaNastroikiArgs) (*CatalogNastroikiRabochegoMestaPolzovateliaNastroiki, error) {
	return r.Client.CatalogNastroikiRabochegoMestaPolzovateliaNastroiki(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiRabochegoMestaPolzovateliaNastroikis(ctx context.Context, args CatalogNastroikiRabochegoMestaPolzovateliaNastroikisArgs) (*[]CatalogNastroikiRabochegoMestaPolzovateliaNastroiki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiRabochegoMestaPolzovateliaNastroikis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiRabochegoMestaPolzovateliaNastroikis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiRabochegoMestaPolzovateliaNastroikis(Where{})
}
func (r *GqlResolver) CatalogsmsShablony(ctx context.Context, args CatalogsmsShablonyArgs) (*CatalogsmsShablony, error) {
	return r.Client.CatalogsmsShablony(args.Key, nil)
}
func (r *GqlResolver) CatalogsmsShablonys(ctx context.Context, args CatalogsmsShablonysArgs) (*[]CatalogsmsShablony, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogsmsShablonys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogsmsShablonys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogsmsShablonys(Where{})
}
func (r *GqlResolver) WriteOff(ctx context.Context, args WriteOffArgs) (*WriteOff, error) {
	return r.Client.WriteOff(args.Key, nil)
}
func (r *GqlResolver) WriteOffs(ctx context.Context, args WriteOffsArgs) (*[]WriteOff, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.WriteOffs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.WriteOffs(Where{Filter: *args.Filter})
	}
	return r.Client.WriteOffs(Where{})
}
func (r *GqlResolver) DocumentSpisanieTovarovTovary(ctx context.Context, args DocumentSpisanieTovarovTovaryArgs) (*DocumentSpisanieTovarovTovary, error) {
	return r.Client.DocumentSpisanieTovarovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentSpisanieTovarovTovarys(ctx context.Context, args DocumentSpisanieTovarovTovarysArgs) (*[]DocumentSpisanieTovarovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSpisanieTovarovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSpisanieTovarovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSpisanieTovarovTovarys(Where{})
}
func (r *GqlResolver) DocumentSpisanieTovarovSertifikaty(ctx context.Context, args DocumentSpisanieTovarovSertifikatyArgs) (*DocumentSpisanieTovarovSertifikaty, error) {
	return r.Client.DocumentSpisanieTovarovSertifikaty(args.Key, nil)
}
func (r *GqlResolver) DocumentSpisanieTovarovSertifikatys(ctx context.Context, args DocumentSpisanieTovarovSertifikatysArgs) (*[]DocumentSpisanieTovarovSertifikaty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSpisanieTovarovSertifikatys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSpisanieTovarovSertifikatys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSpisanieTovarovSertifikatys(Where{})
}
func (r *GqlResolver) DocumentsmsSoobshchenie(ctx context.Context, args DocumentsmsSoobshchenieArgs) (*DocumentsmsSoobshchenie, error) {
	return r.Client.DocumentsmsSoobshchenie(args.Key, nil)
}
func (r *GqlResolver) DocumentsmsSoobshchenies(ctx context.Context, args DocumentsmsSoobshcheniesArgs) (*[]DocumentsmsSoobshchenie, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentsmsSoobshchenies(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentsmsSoobshchenies(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentsmsSoobshchenies(Where{})
}
func (r *GqlResolver) DocumentsmsSoobshcheniePoluchateli(ctx context.Context, args DocumentsmsSoobshcheniePoluchateliArgs) (*DocumentsmsSoobshcheniePoluchateli, error) {
	return r.Client.DocumentsmsSoobshcheniePoluchateli(args.Key, nil)
}
func (r *GqlResolver) DocumentsmsSoobshcheniePoluchatelis(ctx context.Context, args DocumentsmsSoobshcheniePoluchatelisArgs) (*[]DocumentsmsSoobshcheniePoluchateli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentsmsSoobshcheniePoluchatelis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentsmsSoobshcheniePoluchatelis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentsmsSoobshcheniePoluchatelis(Where{})
}
func (r *GqlResolver) DocumentOplataOtPokupateliaPlatezhnoiKartoi(ctx context.Context, args DocumentOplataOtPokupateliaPlatezhnoiKartoiArgs) (*DocumentOplataOtPokupateliaPlatezhnoiKartoi, error) {
	return r.Client.DocumentOplataOtPokupateliaPlatezhnoiKartoi(args.Key, nil)
}
func (r *GqlResolver) DocumentOplataOtPokupateliaPlatezhnoiKartois(ctx context.Context, args DocumentOplataOtPokupateliaPlatezhnoiKartoisArgs) (*[]DocumentOplataOtPokupateliaPlatezhnoiKartoi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOplataOtPokupateliaPlatezhnoiKartois(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOplataOtPokupateliaPlatezhnoiKartois(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOplataOtPokupateliaPlatezhnoiKartois(Where{})
}
func (r *GqlResolver) CatalogDragotsennyeKamni(ctx context.Context, args CatalogDragotsennyeKamniArgs) (*CatalogDragotsennyeKamni, error) {
	return r.Client.CatalogDragotsennyeKamni(args.Key, nil)
}
func (r *GqlResolver) CatalogDragotsennyeKamnis(ctx context.Context, args CatalogDragotsennyeKamnisArgs) (*[]CatalogDragotsennyeKamni, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDragotsennyeKamnis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDragotsennyeKamnis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDragotsennyeKamnis(Where{})
}
func (r *GqlResolver) CatalogKalendariPlanirovaniiaProdazh(ctx context.Context, args CatalogKalendariPlanirovaniiaProdazhArgs) (*CatalogKalendariPlanirovaniiaProdazh, error) {
	return r.Client.CatalogKalendariPlanirovaniiaProdazh(args.Key, nil)
}
func (r *GqlResolver) CatalogKalendariPlanirovaniiaProdazhs(ctx context.Context, args CatalogKalendariPlanirovaniiaProdazhsArgs) (*[]CatalogKalendariPlanirovaniiaProdazh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKalendariPlanirovaniiaProdazhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKalendariPlanirovaniiaProdazhs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKalendariPlanirovaniiaProdazhs(Where{})
}
func (r *GqlResolver) CatalogKalendariPlanirovaniiaProdazhKUPoDniam(ctx context.Context, args CatalogKalendariPlanirovaniiaProdazhKUPoDniamArgs) (*CatalogKalendariPlanirovaniiaProdazhKUPoDniam, error) {
	return r.Client.CatalogKalendariPlanirovaniiaProdazhKUPoDniam(args.Key, nil)
}
func (r *GqlResolver) CatalogKalendariPlanirovaniiaProdazhKUPoDniams(ctx context.Context, args CatalogKalendariPlanirovaniiaProdazhKUPoDniamsArgs) (*[]CatalogKalendariPlanirovaniiaProdazhKUPoDniam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKalendariPlanirovaniiaProdazhKUPoDniams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKalendariPlanirovaniiaProdazhKUPoDniams(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKalendariPlanirovaniiaProdazhKUPoDniams(Where{})
}
func (r *GqlResolver) CatalogKalendariPlanirovaniiaProdazhSalony(ctx context.Context, args CatalogKalendariPlanirovaniiaProdazhSalonyArgs) (*CatalogKalendariPlanirovaniiaProdazhSalony, error) {
	return r.Client.CatalogKalendariPlanirovaniiaProdazhSalony(args.Key, nil)
}
func (r *GqlResolver) CatalogKalendariPlanirovaniiaProdazhSalonys(ctx context.Context, args CatalogKalendariPlanirovaniiaProdazhSalonysArgs) (*[]CatalogKalendariPlanirovaniiaProdazhSalony, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKalendariPlanirovaniiaProdazhSalonys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKalendariPlanirovaniiaProdazhSalonys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKalendariPlanirovaniiaProdazhSalonys(Where{})
}
func (r *GqlResolver) CatalogKontaktnyeLitsa(ctx context.Context, args CatalogKontaktnyeLitsaArgs) (*CatalogKontaktnyeLitsa, error) {
	return r.Client.CatalogKontaktnyeLitsa(args.Key, nil)
}
func (r *GqlResolver) CatalogKontaktnyeLitsas(ctx context.Context, args CatalogKontaktnyeLitsasArgs) (*[]CatalogKontaktnyeLitsa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKontaktnyeLitsas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKontaktnyeLitsas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKontaktnyeLitsas(Where{})
}
func (r *GqlResolver) CatalogFizicheskieLitsa(ctx context.Context, args CatalogFizicheskieLitsaArgs) (*CatalogFizicheskieLitsa, error) {
	return r.Client.CatalogFizicheskieLitsa(args.Key, nil)
}
func (r *GqlResolver) CatalogFizicheskieLitsas(ctx context.Context, args CatalogFizicheskieLitsasArgs) (*[]CatalogFizicheskieLitsa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFizicheskieLitsas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFizicheskieLitsas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFizicheskieLitsas(Where{})
}
func (r *GqlResolver) CatalogTipovyeAnkety(ctx context.Context, args CatalogTipovyeAnketyArgs) (*CatalogTipovyeAnkety, error) {
	return r.Client.CatalogTipovyeAnkety(args.Key, nil)
}
func (r *GqlResolver) CatalogTipovyeAnketys(ctx context.Context, args CatalogTipovyeAnketysArgs) (*[]CatalogTipovyeAnkety, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTipovyeAnketys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTipovyeAnketys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTipovyeAnketys(Where{})
}
func (r *GqlResolver) CatalogTipovyeAnketyVoprosyAnkety(ctx context.Context, args CatalogTipovyeAnketyVoprosyAnketyArgs) (*CatalogTipovyeAnketyVoprosyAnkety, error) {
	return r.Client.CatalogTipovyeAnketyVoprosyAnkety(args.Key, nil)
}
func (r *GqlResolver) CatalogTipovyeAnketyVoprosyAnketys(ctx context.Context, args CatalogTipovyeAnketyVoprosyAnketysArgs) (*[]CatalogTipovyeAnketyVoprosyAnkety, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTipovyeAnketyVoprosyAnketys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTipovyeAnketyVoprosyAnketys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTipovyeAnketyVoprosyAnketys(Where{})
}
func (r *GqlResolver) DocumentNachislenieSpisanieBonusov(ctx context.Context, args DocumentNachislenieSpisanieBonusovArgs) (*DocumentNachislenieSpisanieBonusov, error) {
	return r.Client.DocumentNachislenieSpisanieBonusov(args.Key, nil)
}
func (r *GqlResolver) DocumentNachislenieSpisanieBonusovs(ctx context.Context, args DocumentNachislenieSpisanieBonusovsArgs) (*[]DocumentNachislenieSpisanieBonusov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentNachislenieSpisanieBonusovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentNachislenieSpisanieBonusovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentNachislenieSpisanieBonusovs(Where{})
}
func (r *GqlResolver) DocumentNachislenieSpisanieBonusovDiskontnyeKarty(ctx context.Context, args DocumentNachislenieSpisanieBonusovDiskontnyeKartyArgs) (*DocumentNachislenieSpisanieBonusovDiskontnyeKarty, error) {
	return r.Client.DocumentNachislenieSpisanieBonusovDiskontnyeKarty(args.Key, nil)
}
func (r *GqlResolver) DocumentNachislenieSpisanieBonusovDiskontnyeKartys(ctx context.Context, args DocumentNachislenieSpisanieBonusovDiskontnyeKartysArgs) (*[]DocumentNachislenieSpisanieBonusovDiskontnyeKarty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentNachislenieSpisanieBonusovDiskontnyeKartys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentNachislenieSpisanieBonusovDiskontnyeKartys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentNachislenieSpisanieBonusovDiskontnyeKartys(Where{})
}
func (r *GqlResolver) Type(ctx context.Context, args TypeArgs) (*Type, error) {
	return r.Client.Type(args.Key, nil)
}
func (r *GqlResolver) Types(ctx context.Context, args TypesArgs) (*[]Type, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Types(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Types(Where{Filter: *args.Filter})
	}
	return r.Client.Types(Where{})
}
func (r *GqlResolver) CatalogfmKodyVidovDokumentov(ctx context.Context, args CatalogfmKodyVidovDokumentovArgs) (*CatalogfmKodyVidovDokumentov, error) {
	return r.Client.CatalogfmKodyVidovDokumentov(args.Key, nil)
}
func (r *GqlResolver) CatalogfmKodyVidovDokumentovs(ctx context.Context, args CatalogfmKodyVidovDokumentovsArgs) (*[]CatalogfmKodyVidovDokumentov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogfmKodyVidovDokumentovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogfmKodyVidovDokumentovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogfmKodyVidovDokumentovs(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovaniePoluchennoe(ctx context.Context, args DocumentPlatezhnoeTrebovaniePoluchennoeArgs) (*DocumentPlatezhnoeTrebovaniePoluchennoe, error) {
	return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoe(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovaniePoluchennoes(ctx context.Context, args DocumentPlatezhnoeTrebovaniePoluchennoesArgs) (*[]DocumentPlatezhnoeTrebovaniePoluchennoe, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoes(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoes(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha(ctx context.Context, args DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhaArgs) (*DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha, error) {
	return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhas(ctx context.Context, args DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhasArgs) (*[]DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoeRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta(ctx context.Context, args DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentaArgs) (*DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta, error) {
	return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentas(ctx context.Context, args DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentasArgs) (*[]DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoeTrebovaniePoluchennoeRekvizityKontragentas(Where{})
}
func (r *GqlResolver) DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv(ctx context.Context, args DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvArgs) (*DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv, error) {
	return r.Client.DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvs(ctx context.Context, args DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvsArgs) (*[]DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvs(Where{})
}
func (r *GqlResolver) DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS(ctx context.Context, args DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSArgs) (*DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS, error) {
	return r.Client.DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS(args.Key, nil)
}
func (r *GqlResolver) DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSs(ctx context.Context, args DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSsArgs) (*[]DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDS, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakrytiePlaniruemykhPostupleniiDenezhnykhSredstvPlaniruemyePostupleniiaDSs(Where{})
}
func (r *GqlResolver) CatalogRazdelyAnkety(ctx context.Context, args CatalogRazdelyAnketyArgs) (*CatalogRazdelyAnkety, error) {
	return r.Client.CatalogRazdelyAnkety(args.Key, nil)
}
func (r *GqlResolver) CatalogRazdelyAnketys(ctx context.Context, args CatalogRazdelyAnketysArgs) (*[]CatalogRazdelyAnkety, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogRazdelyAnketys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogRazdelyAnketys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogRazdelyAnketys(Where{})
}
func (r *GqlResolver) DocumentOtchetPoFinMonitoringu(ctx context.Context, args DocumentOtchetPoFinMonitoringuArgs) (*DocumentOtchetPoFinMonitoringu, error) {
	return r.Client.DocumentOtchetPoFinMonitoringu(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetPoFinMonitoringus(ctx context.Context, args DocumentOtchetPoFinMonitoringusArgs) (*[]DocumentOtchetPoFinMonitoringu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetPoFinMonitoringus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetPoFinMonitoringus(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetPoFinMonitoringus(Where{})
}
func (r *GqlResolver) DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa(ctx context.Context, args DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringaArgs) (*DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa, error) {
	return r.Client.DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringas(ctx context.Context, args DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringasArgs) (*[]DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetPoFinMonitoringuDokumentyFinMonitoringas(Where{})
}
func (r *GqlResolver) DocumentOtchetPoFinMonitoringuDannyeDokumenta(ctx context.Context, args DocumentOtchetPoFinMonitoringuDannyeDokumentaArgs) (*DocumentOtchetPoFinMonitoringuDannyeDokumenta, error) {
	return r.Client.DocumentOtchetPoFinMonitoringuDannyeDokumenta(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetPoFinMonitoringuDannyeDokumentas(ctx context.Context, args DocumentOtchetPoFinMonitoringuDannyeDokumentasArgs) (*[]DocumentOtchetPoFinMonitoringuDannyeDokumenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetPoFinMonitoringuDannyeDokumentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetPoFinMonitoringuDannyeDokumentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetPoFinMonitoringuDannyeDokumentas(Where{})
}
func (r *GqlResolver) CatalogKliuchiAnalitikiUchetaNomenklatury(ctx context.Context, args CatalogKliuchiAnalitikiUchetaNomenklaturyArgs) (*CatalogKliuchiAnalitikiUchetaNomenklatury, error) {
	return r.Client.CatalogKliuchiAnalitikiUchetaNomenklatury(args.Key, nil)
}
func (r *GqlResolver) CatalogKliuchiAnalitikiUchetaNomenklaturys(ctx context.Context, args CatalogKliuchiAnalitikiUchetaNomenklaturysArgs) (*[]CatalogKliuchiAnalitikiUchetaNomenklatury, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKliuchiAnalitikiUchetaNomenklaturys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKliuchiAnalitikiUchetaNomenklaturys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKliuchiAnalitikiUchetaNomenklaturys(Where{})
}
func (r *GqlResolver) CatalogVersiiFailov(ctx context.Context, args CatalogVersiiFailovArgs) (*CatalogVersiiFailov, error) {
	return r.Client.CatalogVersiiFailov(args.Key, nil)
}
func (r *GqlResolver) CatalogVersiiFailovs(ctx context.Context, args CatalogVersiiFailovsArgs) (*[]CatalogVersiiFailov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVersiiFailovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVersiiFailovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVersiiFailovs(Where{})
}
func (r *GqlResolver) CatalogVersiiFailovElektronnyePodpisi(ctx context.Context, args CatalogVersiiFailovElektronnyePodpisiArgs) (*CatalogVersiiFailovElektronnyePodpisi, error) {
	return r.Client.CatalogVersiiFailovElektronnyePodpisi(args.Key, nil)
}
func (r *GqlResolver) CatalogVersiiFailovElektronnyePodpisis(ctx context.Context, args CatalogVersiiFailovElektronnyePodpisisArgs) (*[]CatalogVersiiFailovElektronnyePodpisi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVersiiFailovElektronnyePodpisis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVersiiFailovElektronnyePodpisis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVersiiFailovElektronnyePodpisis(Where{})
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklatury(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyArgs) (*DocumentUstanovkaTsenNomenklatury, error) {
	return r.Client.DocumentUstanovkaTsenNomenklatury(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturys(ctx context.Context, args DocumentUstanovkaTsenNomenklaturysArgs) (*[]DocumentUstanovkaTsenNomenklatury, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaTsenNomenklaturys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaTsenNomenklaturys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaTsenNomenklaturys(Where{})
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyTipyTsen(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyTipyTsenArgs) (*DocumentUstanovkaTsenNomenklaturyTipyTsen, error) {
	return r.Client.DocumentUstanovkaTsenNomenklaturyTipyTsen(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyTipyTsens(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyTipyTsensArgs) (*[]DocumentUstanovkaTsenNomenklaturyTipyTsen, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaTsenNomenklaturyTipyTsens(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaTsenNomenklaturyTipyTsens(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaTsenNomenklaturyTipyTsens(Where{})
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyTovary(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyTovaryArgs) (*DocumentUstanovkaTsenNomenklaturyTovary, error) {
	return r.Client.DocumentUstanovkaTsenNomenklaturyTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyTovarys(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyTovarysArgs) (*[]DocumentUstanovkaTsenNomenklaturyTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaTsenNomenklaturyTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaTsenNomenklaturyTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaTsenNomenklaturyTovarys(Where{})
}
func (r *GqlResolver) DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv(ctx context.Context, args DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvArgs) (*DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv, error) {
	return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvs(ctx context.Context, args DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvsArgs) (*[]DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvs(Where{})
}
func (r *GqlResolver) DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha(ctx context.Context, args DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhaArgs) (*DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha, error) {
	return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhas(ctx context.Context, args DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhasArgs) (*[]DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta(ctx context.Context, args DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentaArgs) (*DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta, error) {
	return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentas(ctx context.Context, args DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentasArgs) (*[]DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnyiOrderSpisanieDenezhnykhSredstvRekvizityKontragentas(Where{})
}
func (r *GqlResolver) DocumentPreiskurantNaSkupku(ctx context.Context, args DocumentPreiskurantNaSkupkuArgs) (*DocumentPreiskurantNaSkupku, error) {
	return r.Client.DocumentPreiskurantNaSkupku(args.Key, nil)
}
func (r *GqlResolver) DocumentPreiskurantNaSkupkus(ctx context.Context, args DocumentPreiskurantNaSkupkusArgs) (*[]DocumentPreiskurantNaSkupku, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPreiskurantNaSkupkus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPreiskurantNaSkupkus(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPreiskurantNaSkupkus(Where{})
}
func (r *GqlResolver) DocumentPreiskurantNaSkupkuProby(ctx context.Context, args DocumentPreiskurantNaSkupkuProbyArgs) (*DocumentPreiskurantNaSkupkuProby, error) {
	return r.Client.DocumentPreiskurantNaSkupkuProby(args.Key, nil)
}
func (r *GqlResolver) DocumentPreiskurantNaSkupkuProbys(ctx context.Context, args DocumentPreiskurantNaSkupkuProbysArgs) (*[]DocumentPreiskurantNaSkupkuProby, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPreiskurantNaSkupkuProbys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPreiskurantNaSkupkuProbys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPreiskurantNaSkupkuProbys(Where{})
}
func (r *GqlResolver) DocumentPeredachaMaterialovVProizvodstvo(ctx context.Context, args DocumentPeredachaMaterialovVProizvodstvoArgs) (*DocumentPeredachaMaterialovVProizvodstvo, error) {
	return r.Client.DocumentPeredachaMaterialovVProizvodstvo(args.Key, nil)
}
func (r *GqlResolver) DocumentPeredachaMaterialovVProizvodstvos(ctx context.Context, args DocumentPeredachaMaterialovVProizvodstvosArgs) (*[]DocumentPeredachaMaterialovVProizvodstvo, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPeredachaMaterialovVProizvodstvos(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPeredachaMaterialovVProizvodstvos(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPeredachaMaterialovVProizvodstvos(Where{})
}
func (r *GqlResolver) DocumentPeredachaMaterialovVProizvodstvoTovary(ctx context.Context, args DocumentPeredachaMaterialovVProizvodstvoTovaryArgs) (*DocumentPeredachaMaterialovVProizvodstvoTovary, error) {
	return r.Client.DocumentPeredachaMaterialovVProizvodstvoTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPeredachaMaterialovVProizvodstvoTovarys(ctx context.Context, args DocumentPeredachaMaterialovVProizvodstvoTovarysArgs) (*[]DocumentPeredachaMaterialovVProizvodstvoTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPeredachaMaterialovVProizvodstvoTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPeredachaMaterialovVProizvodstvoTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPeredachaMaterialovVProizvodstvoTovarys(Where{})
}
func (r *GqlResolver) DocumentVnutrenniiZakaz(ctx context.Context, args DocumentVnutrenniiZakazArgs) (*DocumentVnutrenniiZakaz, error) {
	return r.Client.DocumentVnutrenniiZakaz(args.Key, nil)
}
func (r *GqlResolver) DocumentVnutrenniiZakazs(ctx context.Context, args DocumentVnutrenniiZakazsArgs) (*[]DocumentVnutrenniiZakaz, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVnutrenniiZakazs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVnutrenniiZakazs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVnutrenniiZakazs(Where{})
}
func (r *GqlResolver) DocumentVnutrenniiZakazTovary(ctx context.Context, args DocumentVnutrenniiZakazTovaryArgs) (*DocumentVnutrenniiZakazTovary, error) {
	return r.Client.DocumentVnutrenniiZakazTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentVnutrenniiZakazTovarys(ctx context.Context, args DocumentVnutrenniiZakazTovarysArgs) (*[]DocumentVnutrenniiZakazTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVnutrenniiZakazTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVnutrenniiZakazTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVnutrenniiZakazTovarys(Where{})
}
func (r *GqlResolver) CatalogKhranilishcheDopolnitelnoiInformatsii(ctx context.Context, args CatalogKhranilishcheDopolnitelnoiInformatsiiArgs) (*CatalogKhranilishcheDopolnitelnoiInformatsii, error) {
	return r.Client.CatalogKhranilishcheDopolnitelnoiInformatsii(args.Key, nil)
}
func (r *GqlResolver) CatalogKhranilishcheDopolnitelnoiInformatsiis(ctx context.Context, args CatalogKhranilishcheDopolnitelnoiInformatsiisArgs) (*[]CatalogKhranilishcheDopolnitelnoiInformatsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKhranilishcheDopolnitelnoiInformatsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKhranilishcheDopolnitelnoiInformatsiis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKhranilishcheDopolnitelnoiInformatsiis(Where{})
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotki(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiArgs) (*CatalogDopolnitelnyeVneshnieObrabotki, error) {
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotki(args.Key, nil)
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkis(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkisArgs) (*[]CatalogDopolnitelnyeVneshnieObrabotki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkis(Where{})
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnostArgs) (*CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost, error) {
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost(args.Key, nil)
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnosts(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnostsArgs) (*[]CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnost, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnosts(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnosts(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiPrinadlezhnosts(Where{})
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiKomandy(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiKomandyArgs) (*CatalogDopolnitelnyeVneshnieObrabotkiKomandy, error) {
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiKomandy(args.Key, nil)
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiKomandys(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiKomandysArgs) (*[]CatalogDopolnitelnyeVneshnieObrabotkiKomandy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiKomandys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiKomandys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiKomandys(Where{})
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiRazdely(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiRazdelyArgs) (*CatalogDopolnitelnyeVneshnieObrabotkiRazdely, error) {
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiRazdely(args.Key, nil)
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiRazdelys(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiRazdelysArgs) (*[]CatalogDopolnitelnyeVneshnieObrabotkiRazdely, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiRazdelys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiRazdelys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiRazdelys(Where{})
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiNaznachenie(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiNaznachenieArgs) (*CatalogDopolnitelnyeVneshnieObrabotkiNaznachenie, error) {
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiNaznachenie(args.Key, nil)
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiNaznachenies(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiNaznacheniesArgs) (*[]CatalogDopolnitelnyeVneshnieObrabotkiNaznachenie, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiNaznachenies(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiNaznachenies(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiNaznachenies(Where{})
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiRazresheniia(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiRazresheniiaArgs) (*CatalogDopolnitelnyeVneshnieObrabotkiRazresheniia, error) {
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiRazresheniia(args.Key, nil)
}
func (r *GqlResolver) CatalogDopolnitelnyeVneshnieObrabotkiRazresheniias(ctx context.Context, args CatalogDopolnitelnyeVneshnieObrabotkiRazresheniiasArgs) (*[]CatalogDopolnitelnyeVneshnieObrabotkiRazresheniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiRazresheniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiRazresheniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDopolnitelnyeVneshnieObrabotkiRazresheniias(Where{})
}
func (r *GqlResolver) CatalogGruppyPolzovatelei(ctx context.Context, args CatalogGruppyPolzovateleiArgs) (*CatalogGruppyPolzovatelei, error) {
	return r.Client.CatalogGruppyPolzovatelei(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyPolzovateleis(ctx context.Context, args CatalogGruppyPolzovateleisArgs) (*[]CatalogGruppyPolzovatelei, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyPolzovateleis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyPolzovateleis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyPolzovateleis(Where{})
}
func (r *GqlResolver) CatalogGruppyPolzovateleiPolzovateliGruppy(ctx context.Context, args CatalogGruppyPolzovateleiPolzovateliGruppyArgs) (*CatalogGruppyPolzovateleiPolzovateliGruppy, error) {
	return r.Client.CatalogGruppyPolzovateleiPolzovateliGruppy(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyPolzovateleiPolzovateliGruppys(ctx context.Context, args CatalogGruppyPolzovateleiPolzovateliGruppysArgs) (*[]CatalogGruppyPolzovateleiPolzovateliGruppy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyPolzovateleiPolzovateliGruppys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyPolzovateleiPolzovateliGruppys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyPolzovateleiPolzovateliGruppys(Where{})
}
func (r *GqlResolver) DocumentJournalZakazyKlientov(ctx context.Context, args DocumentJournalZakazyKlientovArgs) (*DocumentJournalZakazyKlientov, error) {
	return r.Client.DocumentJournalZakazyKlientov(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalZakazyKlientovs(ctx context.Context, args DocumentJournalZakazyKlientovsArgs) (*[]DocumentJournalZakazyKlientov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalZakazyKlientovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalZakazyKlientovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalZakazyKlientovs(Where{})
}
func (r *GqlResolver) DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki(ctx context.Context, args DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiArgs) (*DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki, error) {
	return r.Client.DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkis(ctx context.Context, args DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkisArgs) (*[]DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkis(Where{})
}
func (r *GqlResolver) DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary(ctx context.Context, args DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovaryArgs) (*DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary, error) {
	return r.Client.DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovarys(ctx context.Context, args DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovarysArgs) (*[]DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratTovarovPostavshchikuIzNeavtomatizirovannoiTorgovoiTochkiTovarys(Where{})
}
func (r *GqlResolver) DocumentZaiavkaNaPeremeshchenieTovarov(ctx context.Context, args DocumentZaiavkaNaPeremeshchenieTovarovArgs) (*DocumentZaiavkaNaPeremeshchenieTovarov, error) {
	return r.Client.DocumentZaiavkaNaPeremeshchenieTovarov(args.Key, nil)
}
func (r *GqlResolver) DocumentZaiavkaNaPeremeshchenieTovarovs(ctx context.Context, args DocumentZaiavkaNaPeremeshchenieTovarovsArgs) (*[]DocumentZaiavkaNaPeremeshchenieTovarov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZaiavkaNaPeremeshchenieTovarovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZaiavkaNaPeremeshchenieTovarovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZaiavkaNaPeremeshchenieTovarovs(Where{})
}
func (r *GqlResolver) DocumentZaiavkaNaPeremeshchenieTovarovTovary(ctx context.Context, args DocumentZaiavkaNaPeremeshchenieTovarovTovaryArgs) (*DocumentZaiavkaNaPeremeshchenieTovarovTovary, error) {
	return r.Client.DocumentZaiavkaNaPeremeshchenieTovarovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentZaiavkaNaPeremeshchenieTovarovTovarys(ctx context.Context, args DocumentZaiavkaNaPeremeshchenieTovarovTovarysArgs) (*[]DocumentZaiavkaNaPeremeshchenieTovarovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZaiavkaNaPeremeshchenieTovarovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZaiavkaNaPeremeshchenieTovarovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZaiavkaNaPeremeshchenieTovarovTovarys(Where{})
}
func (r *GqlResolver) CatalogUsloviiaProdazh(ctx context.Context, args CatalogUsloviiaProdazhArgs) (*CatalogUsloviiaProdazh, error) {
	return r.Client.CatalogUsloviiaProdazh(args.Key, nil)
}
func (r *GqlResolver) CatalogUsloviiaProdazhs(ctx context.Context, args CatalogUsloviiaProdazhsArgs) (*[]CatalogUsloviiaProdazh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUsloviiaProdazhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUsloviiaProdazhs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUsloviiaProdazhs(Where{})
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoFinMonitoringu(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoFinMonitoringuArgs) (*DocumentVvodNachalnykhOstatkovPoFinMonitoringu, error) {
	return r.Client.DocumentVvodNachalnykhOstatkovPoFinMonitoringu(args.Key, nil)
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoFinMonitoringus(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoFinMonitoringusArgs) (*[]DocumentVvodNachalnykhOstatkovPoFinMonitoringu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVvodNachalnykhOstatkovPoFinMonitoringus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVvodNachalnykhOstatkovPoFinMonitoringus(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVvodNachalnykhOstatkovPoFinMonitoringus(Where{})
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoraArgs) (*DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora, error) {
	return r.Client.DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora(args.Key, nil)
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoras(ctx context.Context, args DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovorasArgs) (*[]DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovora, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoras(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoras(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVvodNachalnykhOstatkovPoFinMonitoringuDogovoras(Where{})
}
func (r *GqlResolver) CatalogOrganizatsii(ctx context.Context, args CatalogOrganizatsiiArgs) (*CatalogOrganizatsii, error) {
	return r.Client.CatalogOrganizatsii(args.Key, nil)
}
func (r *GqlResolver) CatalogOrganizatsiis(ctx context.Context, args CatalogOrganizatsiisArgs) (*[]CatalogOrganizatsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogOrganizatsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogOrganizatsiis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogOrganizatsiis(Where{})
}
func (r *GqlResolver) CatalogUsloviiaOplaty(ctx context.Context, args CatalogUsloviiaOplatyArgs) (*CatalogUsloviiaOplaty, error) {
	return r.Client.CatalogUsloviiaOplaty(args.Key, nil)
}
func (r *GqlResolver) CatalogUsloviiaOplatys(ctx context.Context, args CatalogUsloviiaOplatysArgs) (*[]CatalogUsloviiaOplaty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUsloviiaOplatys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUsloviiaOplatys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUsloviiaOplatys(Where{})
}
func (r *GqlResolver) CatalogUsloviiaOplatyTablitsaVyplat(ctx context.Context, args CatalogUsloviiaOplatyTablitsaVyplatArgs) (*CatalogUsloviiaOplatyTablitsaVyplat, error) {
	return r.Client.CatalogUsloviiaOplatyTablitsaVyplat(args.Key, nil)
}
func (r *GqlResolver) CatalogUsloviiaOplatyTablitsaVyplats(ctx context.Context, args CatalogUsloviiaOplatyTablitsaVyplatsArgs) (*[]CatalogUsloviiaOplatyTablitsaVyplat, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUsloviiaOplatyTablitsaVyplats(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUsloviiaOplatyTablitsaVyplats(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUsloviiaOplatyTablitsaVyplats(Where{})
}
func (r *GqlResolver) CatalogKategoriiObieektov(ctx context.Context, args CatalogKategoriiObieektovArgs) (*CatalogKategoriiObieektov, error) {
	return r.Client.CatalogKategoriiObieektov(args.Key, nil)
}
func (r *GqlResolver) CatalogKategoriiObieektovs(ctx context.Context, args CatalogKategoriiObieektovsArgs) (*[]CatalogKategoriiObieektov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKategoriiObieektovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKategoriiObieektovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKategoriiObieektovs(Where{})
}
func (r *GqlResolver) CatalogfmZnacheniiaDliaZapolneniia(ctx context.Context, args CatalogfmZnacheniiaDliaZapolneniiaArgs) (*CatalogfmZnacheniiaDliaZapolneniia, error) {
	return r.Client.CatalogfmZnacheniiaDliaZapolneniia(args.Key, nil)
}
func (r *GqlResolver) CatalogfmZnacheniiaDliaZapolneniias(ctx context.Context, args CatalogfmZnacheniiaDliaZapolneniiasArgs) (*[]CatalogfmZnacheniiaDliaZapolneniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogfmZnacheniiaDliaZapolneniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogfmZnacheniiaDliaZapolneniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogfmZnacheniiaDliaZapolneniias(Where{})
}
func (r *GqlResolver) DocumentUdalitNariadZakaz(ctx context.Context, args DocumentUdalitNariadZakazArgs) (*DocumentUdalitNariadZakaz, error) {
	return r.Client.DocumentUdalitNariadZakaz(args.Key, nil)
}
func (r *GqlResolver) DocumentUdalitNariadZakazs(ctx context.Context, args DocumentUdalitNariadZakazsArgs) (*[]DocumentUdalitNariadZakaz, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUdalitNariadZakazs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUdalitNariadZakazs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUdalitNariadZakazs(Where{})
}
func (r *GqlResolver) DocumentUdalitNariadZakazIzdeliia(ctx context.Context, args DocumentUdalitNariadZakazIzdeliiaArgs) (*DocumentUdalitNariadZakazIzdeliia, error) {
	return r.Client.DocumentUdalitNariadZakazIzdeliia(args.Key, nil)
}
func (r *GqlResolver) DocumentUdalitNariadZakazIzdeliias(ctx context.Context, args DocumentUdalitNariadZakazIzdeliiasArgs) (*[]DocumentUdalitNariadZakazIzdeliia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUdalitNariadZakazIzdeliias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUdalitNariadZakazIzdeliias(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUdalitNariadZakazIzdeliias(Where{})
}
func (r *GqlResolver) DocumentUdalitNariadZakazUslugi(ctx context.Context, args DocumentUdalitNariadZakazUslugiArgs) (*DocumentUdalitNariadZakazUslugi, error) {
	return r.Client.DocumentUdalitNariadZakazUslugi(args.Key, nil)
}
func (r *GqlResolver) DocumentUdalitNariadZakazUslugis(ctx context.Context, args DocumentUdalitNariadZakazUslugisArgs) (*[]DocumentUdalitNariadZakazUslugi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUdalitNariadZakazUslugis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUdalitNariadZakazUslugis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUdalitNariadZakazUslugis(Where{})
}
func (r *GqlResolver) DocumentUdalitNariadZakazSpetsifikatsiia(ctx context.Context, args DocumentUdalitNariadZakazSpetsifikatsiiaArgs) (*DocumentUdalitNariadZakazSpetsifikatsiia, error) {
	return r.Client.DocumentUdalitNariadZakazSpetsifikatsiia(args.Key, nil)
}
func (r *GqlResolver) DocumentUdalitNariadZakazSpetsifikatsiias(ctx context.Context, args DocumentUdalitNariadZakazSpetsifikatsiiasArgs) (*[]DocumentUdalitNariadZakazSpetsifikatsiia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUdalitNariadZakazSpetsifikatsiias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUdalitNariadZakazSpetsifikatsiias(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUdalitNariadZakazSpetsifikatsiias(Where{})
}
func (r *GqlResolver) DocumentUdalitNariadZakazMetally(ctx context.Context, args DocumentUdalitNariadZakazMetallyArgs) (*DocumentUdalitNariadZakazMetally, error) {
	return r.Client.DocumentUdalitNariadZakazMetally(args.Key, nil)
}
func (r *GqlResolver) DocumentUdalitNariadZakazMetallys(ctx context.Context, args DocumentUdalitNariadZakazMetallysArgs) (*[]DocumentUdalitNariadZakazMetally, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUdalitNariadZakazMetallys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUdalitNariadZakazMetallys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUdalitNariadZakazMetallys(Where{})
}
func (r *GqlResolver) DocumentUdalitNariadZakazVstavki(ctx context.Context, args DocumentUdalitNariadZakazVstavkiArgs) (*DocumentUdalitNariadZakazVstavki, error) {
	return r.Client.DocumentUdalitNariadZakazVstavki(args.Key, nil)
}
func (r *GqlResolver) DocumentUdalitNariadZakazVstavkis(ctx context.Context, args DocumentUdalitNariadZakazVstavkisArgs) (*[]DocumentUdalitNariadZakazVstavki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUdalitNariadZakazVstavkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUdalitNariadZakazVstavkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUdalitNariadZakazVstavkis(Where{})
}
func (r *GqlResolver) CatalogBanki(ctx context.Context, args CatalogBankiArgs) (*CatalogBanki, error) {
	return r.Client.CatalogBanki(args.Key, nil)
}
func (r *GqlResolver) CatalogBankis(ctx context.Context, args CatalogBankisArgs) (*[]CatalogBanki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogBankis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogBankis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogBankis(Where{})
}
func (r *GqlResolver) CatalogRoliKontaktnykhLits(ctx context.Context, args CatalogRoliKontaktnykhLitsArgs) (*CatalogRoliKontaktnykhLits, error) {
	return r.Client.CatalogRoliKontaktnykhLits(args.Key, nil)
}
func (r *GqlResolver) CatalogRoliKontaktnykhLitss(ctx context.Context, args CatalogRoliKontaktnykhLitssArgs) (*[]CatalogRoliKontaktnykhLits, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogRoliKontaktnykhLitss(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogRoliKontaktnykhLitss(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogRoliKontaktnykhLitss(Where{})
}
func (r *GqlResolver) DocumentRestrukturizatsiiaDolga(ctx context.Context, args DocumentRestrukturizatsiiaDolgaArgs) (*DocumentRestrukturizatsiiaDolga, error) {
	return r.Client.DocumentRestrukturizatsiiaDolga(args.Key, nil)
}
func (r *GqlResolver) DocumentRestrukturizatsiiaDolgas(ctx context.Context, args DocumentRestrukturizatsiiaDolgasArgs) (*[]DocumentRestrukturizatsiiaDolga, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRestrukturizatsiiaDolgas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRestrukturizatsiiaDolgas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRestrukturizatsiiaDolgas(Where{})
}
func (r *GqlResolver) DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti(ctx context.Context, args DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostiArgs) (*DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti, error) {
	return r.Client.DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti(args.Key, nil)
}
func (r *GqlResolver) DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostis(ctx context.Context, args DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostisArgs) (*[]DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennosti, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRestrukturizatsiiaDolgaRasshifrovkaZadolzhennostis(Where{})
}
func (r *GqlResolver) DocumentAkkreditivPoluchennyi(ctx context.Context, args DocumentAkkreditivPoluchennyiArgs) (*DocumentAkkreditivPoluchennyi, error) {
	return r.Client.DocumentAkkreditivPoluchennyi(args.Key, nil)
}
func (r *GqlResolver) DocumentAkkreditivPoluchennyis(ctx context.Context, args DocumentAkkreditivPoluchennyisArgs) (*[]DocumentAkkreditivPoluchennyi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAkkreditivPoluchennyis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAkkreditivPoluchennyis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAkkreditivPoluchennyis(Where{})
}
func (r *GqlResolver) DocumentAkkreditivPoluchennyiRasshifrovkaPlatezha(ctx context.Context, args DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhaArgs) (*DocumentAkkreditivPoluchennyiRasshifrovkaPlatezha, error) {
	return r.Client.DocumentAkkreditivPoluchennyiRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhas(ctx context.Context, args DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhasArgs) (*[]DocumentAkkreditivPoluchennyiRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAkkreditivPoluchennyiRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentAkkreditivPoluchennyiRekvizityKontragenta(ctx context.Context, args DocumentAkkreditivPoluchennyiRekvizityKontragentaArgs) (*DocumentAkkreditivPoluchennyiRekvizityKontragenta, error) {
	return r.Client.DocumentAkkreditivPoluchennyiRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentAkkreditivPoluchennyiRekvizityKontragentas(ctx context.Context, args DocumentAkkreditivPoluchennyiRekvizityKontragentasArgs) (*[]DocumentAkkreditivPoluchennyiRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAkkreditivPoluchennyiRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAkkreditivPoluchennyiRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAkkreditivPoluchennyiRekvizityKontragentas(Where{})
}
func (r *GqlResolver) DocumentPriemIzRemonta(ctx context.Context, args DocumentPriemIzRemontaArgs) (*DocumentPriemIzRemonta, error) {
	return r.Client.DocumentPriemIzRemonta(args.Key, nil)
}
func (r *GqlResolver) DocumentPriemIzRemontas(ctx context.Context, args DocumentPriemIzRemontasArgs) (*[]DocumentPriemIzRemonta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPriemIzRemontas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPriemIzRemontas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPriemIzRemontas(Where{})
}
func (r *GqlResolver) DocumentPriemIzRemontaIzdeliia(ctx context.Context, args DocumentPriemIzRemontaIzdeliiaArgs) (*DocumentPriemIzRemontaIzdeliia, error) {
	return r.Client.DocumentPriemIzRemontaIzdeliia(args.Key, nil)
}
func (r *GqlResolver) DocumentPriemIzRemontaIzdeliias(ctx context.Context, args DocumentPriemIzRemontaIzdeliiasArgs) (*[]DocumentPriemIzRemontaIzdeliia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPriemIzRemontaIzdeliias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPriemIzRemontaIzdeliias(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPriemIzRemontaIzdeliias(Where{})
}
func (r *GqlResolver) DocumentPriemIzRemontaMaterialy(ctx context.Context, args DocumentPriemIzRemontaMaterialyArgs) (*DocumentPriemIzRemontaMaterialy, error) {
	return r.Client.DocumentPriemIzRemontaMaterialy(args.Key, nil)
}
func (r *GqlResolver) DocumentPriemIzRemontaMaterialys(ctx context.Context, args DocumentPriemIzRemontaMaterialysArgs) (*[]DocumentPriemIzRemontaMaterialy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPriemIzRemontaMaterialys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPriemIzRemontaMaterialys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPriemIzRemontaMaterialys(Where{})
}
func (r *GqlResolver) CatalogTsveta(ctx context.Context, args CatalogTsvetaArgs) (*CatalogTsveta, error) {
	return r.Client.CatalogTsveta(args.Key, nil)
}
func (r *GqlResolver) CatalogTsvetas(ctx context.Context, args CatalogTsvetasArgs) (*[]CatalogTsveta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTsvetas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTsvetas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTsvetas(Where{})
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomissioneraOProdazhakh(ctx context.Context, args DocumentStornirovanieOtchetaKomissioneraOProdazhakhArgs) (*DocumentStornirovanieOtchetaKomissioneraOProdazhakh, error) {
	return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakh(args.Key, nil)
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomissioneraOProdazhakhs(ctx context.Context, args DocumentStornirovanieOtchetaKomissioneraOProdazhakhsArgs) (*[]DocumentStornirovanieOtchetaKomissioneraOProdazhakh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhs(Where{})
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva(ctx context.Context, args DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvaArgs) (*DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva, error) {
	return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva(args.Key, nil)
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvas(ctx context.Context, args DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvasArgs) (*[]DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstva, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhDenezhnyeSredstvas(Where{})
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary(ctx context.Context, args DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovaryArgs) (*DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary, error) {
	return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovarys(ctx context.Context, args DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovarysArgs) (*[]DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentStornirovanieOtchetaKomissioneraOProdazhakhTovarys(Where{})
}
func (r *GqlResolver) DocumentJournalDavalcheskieDokumenty(ctx context.Context, args DocumentJournalDavalcheskieDokumentyArgs) (*DocumentJournalDavalcheskieDokumenty, error) {
	return r.Client.DocumentJournalDavalcheskieDokumenty(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalDavalcheskieDokumentys(ctx context.Context, args DocumentJournalDavalcheskieDokumentysArgs) (*[]DocumentJournalDavalcheskieDokumenty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalDavalcheskieDokumentys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalDavalcheskieDokumentys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalDavalcheskieDokumentys(Where{})
}
func (r *GqlResolver) CatalogfmAnketaKlienta(ctx context.Context, args CatalogfmAnketaKlientaArgs) (*CatalogfmAnketaKlienta, error) {
	return r.Client.CatalogfmAnketaKlienta(args.Key, nil)
}
func (r *GqlResolver) CatalogfmAnketaKlientas(ctx context.Context, args CatalogfmAnketaKlientasArgs) (*[]CatalogfmAnketaKlienta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogfmAnketaKlientas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogfmAnketaKlientas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogfmAnketaKlientas(Where{})
}
func (r *GqlResolver) CatalogfmAnketaKlientaDannyeKontragenta(ctx context.Context, args CatalogfmAnketaKlientaDannyeKontragentaArgs) (*CatalogfmAnketaKlientaDannyeKontragenta, error) {
	return r.Client.CatalogfmAnketaKlientaDannyeKontragenta(args.Key, nil)
}
func (r *GqlResolver) CatalogfmAnketaKlientaDannyeKontragentas(ctx context.Context, args CatalogfmAnketaKlientaDannyeKontragentasArgs) (*[]CatalogfmAnketaKlientaDannyeKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogfmAnketaKlientaDannyeKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogfmAnketaKlientaDannyeKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogfmAnketaKlientaDannyeKontragentas(Where{})
}
func (r *GqlResolver) CatalogVidyVzaimoraschetov(ctx context.Context, args CatalogVidyVzaimoraschetovArgs) (*CatalogVidyVzaimoraschetov, error) {
	return r.Client.CatalogVidyVzaimoraschetov(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyVzaimoraschetovs(ctx context.Context, args CatalogVidyVzaimoraschetovsArgs) (*[]CatalogVidyVzaimoraschetov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyVzaimoraschetovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyVzaimoraschetovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyVzaimoraschetovs(Where{})
}
func (r *GqlResolver) DocumentUstanovkaZnacheniiTochkiZakaza(ctx context.Context, args DocumentUstanovkaZnacheniiTochkiZakazaArgs) (*DocumentUstanovkaZnacheniiTochkiZakaza, error) {
	return r.Client.DocumentUstanovkaZnacheniiTochkiZakaza(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaZnacheniiTochkiZakazas(ctx context.Context, args DocumentUstanovkaZnacheniiTochkiZakazasArgs) (*[]DocumentUstanovkaZnacheniiTochkiZakaza, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaZnacheniiTochkiZakazas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaZnacheniiTochkiZakazas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaZnacheniiTochkiZakazas(Where{})
}
func (r *GqlResolver) DocumentUstanovkaZnacheniiTochkiZakazaTovary(ctx context.Context, args DocumentUstanovkaZnacheniiTochkiZakazaTovaryArgs) (*DocumentUstanovkaZnacheniiTochkiZakazaTovary, error) {
	return r.Client.DocumentUstanovkaZnacheniiTochkiZakazaTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaZnacheniiTochkiZakazaTovarys(ctx context.Context, args DocumentUstanovkaZnacheniiTochkiZakazaTovarysArgs) (*[]DocumentUstanovkaZnacheniiTochkiZakazaTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaZnacheniiTochkiZakazaTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaZnacheniiTochkiZakazaTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaZnacheniiTochkiZakazaTovarys(Where{})
}
func (r *GqlResolver) CatalogZnacheniiaKodirovki(ctx context.Context, args CatalogZnacheniiaKodirovkiArgs) (*CatalogZnacheniiaKodirovki, error) {
	return r.Client.CatalogZnacheniiaKodirovki(args.Key, nil)
}
func (r *GqlResolver) CatalogZnacheniiaKodirovkis(ctx context.Context, args CatalogZnacheniiaKodirovkisArgs) (*[]CatalogZnacheniiaKodirovki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogZnacheniiaKodirovkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogZnacheniiaKodirovkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogZnacheniiaKodirovkis(Where{})
}
func (r *GqlResolver) CatalogPravilaProdazh(ctx context.Context, args CatalogPravilaProdazhArgs) (*CatalogPravilaProdazh, error) {
	return r.Client.CatalogPravilaProdazh(args.Key, nil)
}
func (r *GqlResolver) CatalogPravilaProdazhs(ctx context.Context, args CatalogPravilaProdazhsArgs) (*[]CatalogPravilaProdazh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogPravilaProdazhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogPravilaProdazhs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogPravilaProdazhs(Where{})
}
func (r *GqlResolver) CatalogPravilaProdazhTovary(ctx context.Context, args CatalogPravilaProdazhTovaryArgs) (*CatalogPravilaProdazhTovary, error) {
	return r.Client.CatalogPravilaProdazhTovary(args.Key, nil)
}
func (r *GqlResolver) CatalogPravilaProdazhTovarys(ctx context.Context, args CatalogPravilaProdazhTovarysArgs) (*[]CatalogPravilaProdazhTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogPravilaProdazhTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogPravilaProdazhTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogPravilaProdazhTovarys(Where{})
}
func (r *GqlResolver) DocumentPostuplenieDopRaskhodov(ctx context.Context, args DocumentPostuplenieDopRaskhodovArgs) (*DocumentPostuplenieDopRaskhodov, error) {
	return r.Client.DocumentPostuplenieDopRaskhodov(args.Key, nil)
}
func (r *GqlResolver) DocumentPostuplenieDopRaskhodovs(ctx context.Context, args DocumentPostuplenieDopRaskhodovsArgs) (*[]DocumentPostuplenieDopRaskhodov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPostuplenieDopRaskhodovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPostuplenieDopRaskhodovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPostuplenieDopRaskhodovs(Where{})
}
func (r *GqlResolver) DocumentPostuplenieDopRaskhodovTovary(ctx context.Context, args DocumentPostuplenieDopRaskhodovTovaryArgs) (*DocumentPostuplenieDopRaskhodovTovary, error) {
	return r.Client.DocumentPostuplenieDopRaskhodovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPostuplenieDopRaskhodovTovarys(ctx context.Context, args DocumentPostuplenieDopRaskhodovTovarysArgs) (*[]DocumentPostuplenieDopRaskhodovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPostuplenieDopRaskhodovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPostuplenieDopRaskhodovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPostuplenieDopRaskhodovTovarys(Where{})
}
func (r *GqlResolver) CatalogKhoziaistvennyeOperatsii(ctx context.Context, args CatalogKhoziaistvennyeOperatsiiArgs) (*CatalogKhoziaistvennyeOperatsii, error) {
	return r.Client.CatalogKhoziaistvennyeOperatsii(args.Key, nil)
}
func (r *GqlResolver) CatalogKhoziaistvennyeOperatsiis(ctx context.Context, args CatalogKhoziaistvennyeOperatsiisArgs) (*[]CatalogKhoziaistvennyeOperatsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKhoziaistvennyeOperatsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKhoziaistvennyeOperatsiis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKhoziaistvennyeOperatsiis(Where{})
}
func (r *GqlResolver) DocumentAvansovyiOtchet(ctx context.Context, args DocumentAvansovyiOtchetArgs) (*DocumentAvansovyiOtchet, error) {
	return r.Client.DocumentAvansovyiOtchet(args.Key, nil)
}
func (r *GqlResolver) DocumentAvansovyiOtchets(ctx context.Context, args DocumentAvansovyiOtchetsArgs) (*[]DocumentAvansovyiOtchet, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAvansovyiOtchets(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAvansovyiOtchets(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAvansovyiOtchets(Where{})
}
func (r *GqlResolver) DocumentAvansovyiOtchetVydannyeAvansy(ctx context.Context, args DocumentAvansovyiOtchetVydannyeAvansyArgs) (*DocumentAvansovyiOtchetVydannyeAvansy, error) {
	return r.Client.DocumentAvansovyiOtchetVydannyeAvansy(args.Key, nil)
}
func (r *GqlResolver) DocumentAvansovyiOtchetVydannyeAvansys(ctx context.Context, args DocumentAvansovyiOtchetVydannyeAvansysArgs) (*[]DocumentAvansovyiOtchetVydannyeAvansy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAvansovyiOtchetVydannyeAvansys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAvansovyiOtchetVydannyeAvansys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAvansovyiOtchetVydannyeAvansys(Where{})
}
func (r *GqlResolver) DocumentAvansovyiOtchetTovary(ctx context.Context, args DocumentAvansovyiOtchetTovaryArgs) (*DocumentAvansovyiOtchetTovary, error) {
	return r.Client.DocumentAvansovyiOtchetTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentAvansovyiOtchetTovarys(ctx context.Context, args DocumentAvansovyiOtchetTovarysArgs) (*[]DocumentAvansovyiOtchetTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAvansovyiOtchetTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAvansovyiOtchetTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAvansovyiOtchetTovarys(Where{})
}
func (r *GqlResolver) DocumentAvansovyiOtchetOplataPostavshchikam(ctx context.Context, args DocumentAvansovyiOtchetOplataPostavshchikamArgs) (*DocumentAvansovyiOtchetOplataPostavshchikam, error) {
	return r.Client.DocumentAvansovyiOtchetOplataPostavshchikam(args.Key, nil)
}
func (r *GqlResolver) DocumentAvansovyiOtchetOplataPostavshchikams(ctx context.Context, args DocumentAvansovyiOtchetOplataPostavshchikamsArgs) (*[]DocumentAvansovyiOtchetOplataPostavshchikam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAvansovyiOtchetOplataPostavshchikams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAvansovyiOtchetOplataPostavshchikams(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAvansovyiOtchetOplataPostavshchikams(Where{})
}
func (r *GqlResolver) DocumentAvansovyiOtchetProchee(ctx context.Context, args DocumentAvansovyiOtchetProcheeArgs) (*DocumentAvansovyiOtchetProchee, error) {
	return r.Client.DocumentAvansovyiOtchetProchee(args.Key, nil)
}
func (r *GqlResolver) DocumentAvansovyiOtchetProchees(ctx context.Context, args DocumentAvansovyiOtchetProcheesArgs) (*[]DocumentAvansovyiOtchetProchee, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentAvansovyiOtchetProchees(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentAvansovyiOtchetProchees(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentAvansovyiOtchetProchees(Where{})
}
func (r *GqlResolver) CatalogDolzhnostiOrganizatsii(ctx context.Context, args CatalogDolzhnostiOrganizatsiiArgs) (*CatalogDolzhnostiOrganizatsii, error) {
	return r.Client.CatalogDolzhnostiOrganizatsii(args.Key, nil)
}
func (r *GqlResolver) CatalogDolzhnostiOrganizatsiis(ctx context.Context, args CatalogDolzhnostiOrganizatsiisArgs) (*[]CatalogDolzhnostiOrganizatsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDolzhnostiOrganizatsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDolzhnostiOrganizatsiis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDolzhnostiOrganizatsiis(Where{})
}
func (r *GqlResolver) CatalogAnalitikaTipaIzdeliia(ctx context.Context, args CatalogAnalitikaTipaIzdeliiaArgs) (*CatalogAnalitikaTipaIzdeliia, error) {
	return r.Client.CatalogAnalitikaTipaIzdeliia(args.Key, nil)
}
func (r *GqlResolver) CatalogAnalitikaTipaIzdeliias(ctx context.Context, args CatalogAnalitikaTipaIzdeliiasArgs) (*[]CatalogAnalitikaTipaIzdeliia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogAnalitikaTipaIzdeliias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogAnalitikaTipaIzdeliias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogAnalitikaTipaIzdeliias(Where{})
}
func (r *GqlResolver) CatalogDopolnitelnyePechatnyeFormy(ctx context.Context, args CatalogDopolnitelnyePechatnyeFormyArgs) (*CatalogDopolnitelnyePechatnyeFormy, error) {
	return r.Client.CatalogDopolnitelnyePechatnyeFormy(args.Key, nil)
}
func (r *GqlResolver) CatalogDopolnitelnyePechatnyeFormys(ctx context.Context, args CatalogDopolnitelnyePechatnyeFormysArgs) (*[]CatalogDopolnitelnyePechatnyeFormy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDopolnitelnyePechatnyeFormys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDopolnitelnyePechatnyeFormys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDopolnitelnyePechatnyeFormys(Where{})
}
func (r *GqlResolver) CatalogDopolnitelnyePechatnyeFormyPrinadlezhnost(ctx context.Context, args CatalogDopolnitelnyePechatnyeFormyPrinadlezhnostArgs) (*CatalogDopolnitelnyePechatnyeFormyPrinadlezhnost, error) {
	return r.Client.CatalogDopolnitelnyePechatnyeFormyPrinadlezhnost(args.Key, nil)
}
func (r *GqlResolver) CatalogDopolnitelnyePechatnyeFormyPrinadlezhnosts(ctx context.Context, args CatalogDopolnitelnyePechatnyeFormyPrinadlezhnostsArgs) (*[]CatalogDopolnitelnyePechatnyeFormyPrinadlezhnost, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDopolnitelnyePechatnyeFormyPrinadlezhnosts(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDopolnitelnyePechatnyeFormyPrinadlezhnosts(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDopolnitelnyePechatnyeFormyPrinadlezhnosts(Where{})
}
func (r *GqlResolver) MemberCardsType(ctx context.Context, args MemberCardsTypeArgs) (*MemberCardsType, error) {
	return r.Client.MemberCardsType(args.Key, nil)
}
func (r *GqlResolver) MemberCardsTypes(ctx context.Context, args MemberCardsTypesArgs) (*[]MemberCardsType, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.MemberCardsTypes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.MemberCardsTypes(Where{Filter: *args.Filter})
	}
	return r.Client.MemberCardsTypes(Where{})
}
func (r *GqlResolver) DocumentRegistratsiiaNaSaite(ctx context.Context, args DocumentRegistratsiiaNaSaiteArgs) (*DocumentRegistratsiiaNaSaite, error) {
	return r.Client.DocumentRegistratsiiaNaSaite(args.Key, nil)
}
func (r *GqlResolver) DocumentRegistratsiiaNaSaites(ctx context.Context, args DocumentRegistratsiiaNaSaitesArgs) (*[]DocumentRegistratsiiaNaSaite, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRegistratsiiaNaSaites(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRegistratsiiaNaSaites(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRegistratsiiaNaSaites(Where{})
}
func (r *GqlResolver) CatalogObrabotkiObsluzhivaniiaTO(ctx context.Context, args CatalogObrabotkiObsluzhivaniiaTOArgs) (*CatalogObrabotkiObsluzhivaniiaTO, error) {
	return r.Client.CatalogObrabotkiObsluzhivaniiaTO(args.Key, nil)
}
func (r *GqlResolver) CatalogObrabotkiObsluzhivaniiaTOs(ctx context.Context, args CatalogObrabotkiObsluzhivaniiaTOsArgs) (*[]CatalogObrabotkiObsluzhivaniiaTO, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogObrabotkiObsluzhivaniiaTOs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogObrabotkiObsluzhivaniiaTOs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogObrabotkiObsluzhivaniiaTOs(Where{})
}
func (r *GqlResolver) CatalogObrabotkiObsluzhivaniiaTOModeli(ctx context.Context, args CatalogObrabotkiObsluzhivaniiaTOModeliArgs) (*CatalogObrabotkiObsluzhivaniiaTOModeli, error) {
	return r.Client.CatalogObrabotkiObsluzhivaniiaTOModeli(args.Key, nil)
}
func (r *GqlResolver) CatalogObrabotkiObsluzhivaniiaTOModelis(ctx context.Context, args CatalogObrabotkiObsluzhivaniiaTOModelisArgs) (*[]CatalogObrabotkiObsluzhivaniiaTOModeli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogObrabotkiObsluzhivaniiaTOModelis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogObrabotkiObsluzhivaniiaTOModelis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogObrabotkiObsluzhivaniiaTOModelis(Where{})
}
func (r *GqlResolver) CatalogNastroikaIntervalov(ctx context.Context, args CatalogNastroikaIntervalovArgs) (*CatalogNastroikaIntervalov, error) {
	return r.Client.CatalogNastroikaIntervalov(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikaIntervalovs(ctx context.Context, args CatalogNastroikaIntervalovsArgs) (*[]CatalogNastroikaIntervalov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikaIntervalovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikaIntervalovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikaIntervalovs(Where{})
}
func (r *GqlResolver) CatalogNastroikaIntervalovTablichnaiaChast(ctx context.Context, args CatalogNastroikaIntervalovTablichnaiaChastArgs) (*CatalogNastroikaIntervalovTablichnaiaChast, error) {
	return r.Client.CatalogNastroikaIntervalovTablichnaiaChast(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikaIntervalovTablichnaiaChasts(ctx context.Context, args CatalogNastroikaIntervalovTablichnaiaChastsArgs) (*[]CatalogNastroikaIntervalovTablichnaiaChast, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikaIntervalovTablichnaiaChasts(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikaIntervalovTablichnaiaChasts(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikaIntervalovTablichnaiaChasts(Where{})
}
func (r *GqlResolver) CatalogProfiliGruppDostupa(ctx context.Context, args CatalogProfiliGruppDostupaArgs) (*CatalogProfiliGruppDostupa, error) {
	return r.Client.CatalogProfiliGruppDostupa(args.Key, nil)
}
func (r *GqlResolver) CatalogProfiliGruppDostupas(ctx context.Context, args CatalogProfiliGruppDostupasArgs) (*[]CatalogProfiliGruppDostupa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogProfiliGruppDostupas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogProfiliGruppDostupas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogProfiliGruppDostupas(Where{})
}
func (r *GqlResolver) CatalogProfiliGruppDostupaRoli(ctx context.Context, args CatalogProfiliGruppDostupaRoliArgs) (*CatalogProfiliGruppDostupaRoli, error) {
	return r.Client.CatalogProfiliGruppDostupaRoli(args.Key, nil)
}
func (r *GqlResolver) CatalogProfiliGruppDostupaRolis(ctx context.Context, args CatalogProfiliGruppDostupaRolisArgs) (*[]CatalogProfiliGruppDostupaRoli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogProfiliGruppDostupaRolis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogProfiliGruppDostupaRolis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogProfiliGruppDostupaRolis(Where{})
}
func (r *GqlResolver) CatalogProfiliGruppDostupaVidyDostupa(ctx context.Context, args CatalogProfiliGruppDostupaVidyDostupaArgs) (*CatalogProfiliGruppDostupaVidyDostupa, error) {
	return r.Client.CatalogProfiliGruppDostupaVidyDostupa(args.Key, nil)
}
func (r *GqlResolver) CatalogProfiliGruppDostupaVidyDostupas(ctx context.Context, args CatalogProfiliGruppDostupaVidyDostupasArgs) (*[]CatalogProfiliGruppDostupaVidyDostupa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogProfiliGruppDostupaVidyDostupas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogProfiliGruppDostupaVidyDostupas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogProfiliGruppDostupaVidyDostupas(Where{})
}
func (r *GqlResolver) CatalogProfiliGruppDostupaZnacheniiaDostupa(ctx context.Context, args CatalogProfiliGruppDostupaZnacheniiaDostupaArgs) (*CatalogProfiliGruppDostupaZnacheniiaDostupa, error) {
	return r.Client.CatalogProfiliGruppDostupaZnacheniiaDostupa(args.Key, nil)
}
func (r *GqlResolver) CatalogProfiliGruppDostupaZnacheniiaDostupas(ctx context.Context, args CatalogProfiliGruppDostupaZnacheniiaDostupasArgs) (*[]CatalogProfiliGruppDostupaZnacheniiaDostupa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogProfiliGruppDostupaZnacheniiaDostupas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogProfiliGruppDostupaZnacheniiaDostupas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogProfiliGruppDostupaZnacheniiaDostupas(Where{})
}
func (r *GqlResolver) CatalogProfiliGruppDostupaDostupPoPodsistemam(ctx context.Context, args CatalogProfiliGruppDostupaDostupPoPodsistemamArgs) (*CatalogProfiliGruppDostupaDostupPoPodsistemam, error) {
	return r.Client.CatalogProfiliGruppDostupaDostupPoPodsistemam(args.Key, nil)
}
func (r *GqlResolver) CatalogProfiliGruppDostupaDostupPoPodsistemams(ctx context.Context, args CatalogProfiliGruppDostupaDostupPoPodsistemamsArgs) (*[]CatalogProfiliGruppDostupaDostupPoPodsistemam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogProfiliGruppDostupaDostupPoPodsistemams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogProfiliGruppDostupaDostupPoPodsistemams(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogProfiliGruppDostupaDostupPoPodsistemams(Where{})
}
func (r *GqlResolver) CatalogNastroikiDliaKurera(ctx context.Context, args CatalogNastroikiDliaKureraArgs) (*CatalogNastroikiDliaKurera, error) {
	return r.Client.CatalogNastroikiDliaKurera(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiDliaKureras(ctx context.Context, args CatalogNastroikiDliaKurerasArgs) (*[]CatalogNastroikiDliaKurera, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiDliaKureras(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiDliaKureras(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiDliaKureras(Where{})
}
func (r *GqlResolver) CatalogNastroikiDliaKureraSostavNaimenovaniia(ctx context.Context, args CatalogNastroikiDliaKureraSostavNaimenovaniiaArgs) (*CatalogNastroikiDliaKureraSostavNaimenovaniia, error) {
	return r.Client.CatalogNastroikiDliaKureraSostavNaimenovaniia(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiDliaKureraSostavNaimenovaniias(ctx context.Context, args CatalogNastroikiDliaKureraSostavNaimenovaniiasArgs) (*[]CatalogNastroikiDliaKureraSostavNaimenovaniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiDliaKureraSostavNaimenovaniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiDliaKureraSostavNaimenovaniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiDliaKureraSostavNaimenovaniias(Where{})
}
func (r *GqlResolver) CatalogTipyTsenNomenklaturyKontragentov(ctx context.Context, args CatalogTipyTsenNomenklaturyKontragentovArgs) (*CatalogTipyTsenNomenklaturyKontragentov, error) {
	return r.Client.CatalogTipyTsenNomenklaturyKontragentov(args.Key, nil)
}
func (r *GqlResolver) CatalogTipyTsenNomenklaturyKontragentovs(ctx context.Context, args CatalogTipyTsenNomenklaturyKontragentovsArgs) (*[]CatalogTipyTsenNomenklaturyKontragentov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTipyTsenNomenklaturyKontragentovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTipyTsenNomenklaturyKontragentovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTipyTsenNomenklaturyKontragentovs(Where{})
}
func (r *GqlResolver) DocumentJournalTsenoobrazovanie(ctx context.Context, args DocumentJournalTsenoobrazovanieArgs) (*DocumentJournalTsenoobrazovanie, error) {
	return r.Client.DocumentJournalTsenoobrazovanie(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalTsenoobrazovanies(ctx context.Context, args DocumentJournalTsenoobrazovaniesArgs) (*[]DocumentJournalTsenoobrazovanie, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalTsenoobrazovanies(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalTsenoobrazovanies(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalTsenoobrazovanies(Where{})
}
func (r *GqlResolver) CatalogEdinitsyIzmereniia(ctx context.Context, args CatalogEdinitsyIzmereniiaArgs) (*CatalogEdinitsyIzmereniia, error) {
	return r.Client.CatalogEdinitsyIzmereniia(args.Key, nil)
}
func (r *GqlResolver) CatalogEdinitsyIzmereniias(ctx context.Context, args CatalogEdinitsyIzmereniiasArgs) (*[]CatalogEdinitsyIzmereniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogEdinitsyIzmereniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogEdinitsyIzmereniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogEdinitsyIzmereniias(Where{})
}
func (r *GqlResolver) CatalogStatiDvizheniiaDenezhnykhSredstv(ctx context.Context, args CatalogStatiDvizheniiaDenezhnykhSredstvArgs) (*CatalogStatiDvizheniiaDenezhnykhSredstv, error) {
	return r.Client.CatalogStatiDvizheniiaDenezhnykhSredstv(args.Key, nil)
}
func (r *GqlResolver) CatalogStatiDvizheniiaDenezhnykhSredstvs(ctx context.Context, args CatalogStatiDvizheniiaDenezhnykhSredstvsArgs) (*[]CatalogStatiDvizheniiaDenezhnykhSredstv, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogStatiDvizheniiaDenezhnykhSredstvs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogStatiDvizheniiaDenezhnykhSredstvs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogStatiDvizheniiaDenezhnykhSredstvs(Where{})
}
func (r *GqlResolver) DocumentInkassovoePorucheniePoluchennoe(ctx context.Context, args DocumentInkassovoePorucheniePoluchennoeArgs) (*DocumentInkassovoePorucheniePoluchennoe, error) {
	return r.Client.DocumentInkassovoePorucheniePoluchennoe(args.Key, nil)
}
func (r *GqlResolver) DocumentInkassovoePorucheniePoluchennoes(ctx context.Context, args DocumentInkassovoePorucheniePoluchennoesArgs) (*[]DocumentInkassovoePorucheniePoluchennoe, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInkassovoePorucheniePoluchennoes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInkassovoePorucheniePoluchennoes(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInkassovoePorucheniePoluchennoes(Where{})
}
func (r *GqlResolver) DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha(ctx context.Context, args DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhaArgs) (*DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha, error) {
	return r.Client.DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhas(ctx context.Context, args DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhasArgs) (*[]DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInkassovoePorucheniePoluchennoeRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta(ctx context.Context, args DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentaArgs) (*DocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta, error) {
	return r.Client.DocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentas(ctx context.Context, args DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentasArgs) (*[]DocumentInkassovoePorucheniePoluchennoeRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInkassovoePorucheniePoluchennoeRekvizityKontragentas(Where{})
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymiShtrikhM(ctx context.Context, args CatalogNastroikiObmenaDannymiShtrikhMArgs) (*CatalogNastroikiObmenaDannymiShtrikhM, error) {
	return r.Client.CatalogNastroikiObmenaDannymiShtrikhM(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiObmenaDannymiShtrikhMs(ctx context.Context, args CatalogNastroikiObmenaDannymiShtrikhMsArgs) (*[]CatalogNastroikiObmenaDannymiShtrikhM, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiObmenaDannymiShtrikhMs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiObmenaDannymiShtrikhMs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiObmenaDannymiShtrikhMs(Where{})
}
func (r *GqlResolver) CatalogStatiZatrat(ctx context.Context, args CatalogStatiZatratArgs) (*CatalogStatiZatrat, error) {
	return r.Client.CatalogStatiZatrat(args.Key, nil)
}
func (r *GqlResolver) CatalogStatiZatrats(ctx context.Context, args CatalogStatiZatratsArgs) (*[]CatalogStatiZatrat, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogStatiZatrats(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogStatiZatrats(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogStatiZatrats(Where{})
}
func (r *GqlResolver) DocumentVozvratTovarovOtPokupatelia(ctx context.Context, args DocumentVozvratTovarovOtPokupateliaArgs) (*DocumentVozvratTovarovOtPokupatelia, error) {
	return r.Client.DocumentVozvratTovarovOtPokupatelia(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratTovarovOtPokupatelias(ctx context.Context, args DocumentVozvratTovarovOtPokupateliasArgs) (*[]DocumentVozvratTovarovOtPokupatelia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratTovarovOtPokupatelias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratTovarovOtPokupatelias(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratTovarovOtPokupatelias(Where{})
}
func (r *GqlResolver) DocumentVozvratTovarovOtPokupateliaTovary(ctx context.Context, args DocumentVozvratTovarovOtPokupateliaTovaryArgs) (*DocumentVozvratTovarovOtPokupateliaTovary, error) {
	return r.Client.DocumentVozvratTovarovOtPokupateliaTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratTovarovOtPokupateliaTovarys(ctx context.Context, args DocumentVozvratTovarovOtPokupateliaTovarysArgs) (*[]DocumentVozvratTovarovOtPokupateliaTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratTovarovOtPokupateliaTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratTovarovOtPokupateliaTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratTovarovOtPokupateliaTovarys(Where{})
}
func (r *GqlResolver) DocumentVozvratTovarovOtPokupateliaUslugi(ctx context.Context, args DocumentVozvratTovarovOtPokupateliaUslugiArgs) (*DocumentVozvratTovarovOtPokupateliaUslugi, error) {
	return r.Client.DocumentVozvratTovarovOtPokupateliaUslugi(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratTovarovOtPokupateliaUslugis(ctx context.Context, args DocumentVozvratTovarovOtPokupateliaUslugisArgs) (*[]DocumentVozvratTovarovOtPokupateliaUslugi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratTovarovOtPokupateliaUslugis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratTovarovOtPokupateliaUslugis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratTovarovOtPokupateliaUslugis(Where{})
}
func (r *GqlResolver) DocumentZakazPostavshchiku(ctx context.Context, args DocumentZakazPostavshchikuArgs) (*DocumentZakazPostavshchiku, error) {
	return r.Client.DocumentZakazPostavshchiku(args.Key, nil)
}
func (r *GqlResolver) DocumentZakazPostavshchikus(ctx context.Context, args DocumentZakazPostavshchikusArgs) (*[]DocumentZakazPostavshchiku, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakazPostavshchikus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakazPostavshchikus(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakazPostavshchikus(Where{})
}
func (r *GqlResolver) DocumentZakazPostavshchikuTovary(ctx context.Context, args DocumentZakazPostavshchikuTovaryArgs) (*DocumentZakazPostavshchikuTovary, error) {
	return r.Client.DocumentZakazPostavshchikuTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentZakazPostavshchikuTovarys(ctx context.Context, args DocumentZakazPostavshchikuTovarysArgs) (*[]DocumentZakazPostavshchikuTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakazPostavshchikuTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakazPostavshchikuTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakazPostavshchikuTovarys(Where{})
}
func (r *GqlResolver) CatalogSkidkiNatsenki(ctx context.Context, args CatalogSkidkiNatsenkiArgs) (*CatalogSkidkiNatsenki, error) {
	return r.Client.CatalogSkidkiNatsenki(args.Key, nil)
}
func (r *GqlResolver) CatalogSkidkiNatsenkis(ctx context.Context, args CatalogSkidkiNatsenkisArgs) (*[]CatalogSkidkiNatsenki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSkidkiNatsenkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSkidkiNatsenkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSkidkiNatsenkis(Where{})
}
func (r *GqlResolver) CatalogSkidkiNatsenkiUsloviiaPredostavleniia(ctx context.Context, args CatalogSkidkiNatsenkiUsloviiaPredostavleniiaArgs) (*CatalogSkidkiNatsenkiUsloviiaPredostavleniia, error) {
	return r.Client.CatalogSkidkiNatsenkiUsloviiaPredostavleniia(args.Key, nil)
}
func (r *GqlResolver) CatalogSkidkiNatsenkiUsloviiaPredostavleniias(ctx context.Context, args CatalogSkidkiNatsenkiUsloviiaPredostavleniiasArgs) (*[]CatalogSkidkiNatsenkiUsloviiaPredostavleniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSkidkiNatsenkiUsloviiaPredostavleniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSkidkiNatsenkiUsloviiaPredostavleniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSkidkiNatsenkiUsloviiaPredostavleniias(Where{})
}
func (r *GqlResolver) CatalogSkidkiNatsenkiTsenovyeGruppy(ctx context.Context, args CatalogSkidkiNatsenkiTsenovyeGruppyArgs) (*CatalogSkidkiNatsenkiTsenovyeGruppy, error) {
	return r.Client.CatalogSkidkiNatsenkiTsenovyeGruppy(args.Key, nil)
}
func (r *GqlResolver) CatalogSkidkiNatsenkiTsenovyeGruppys(ctx context.Context, args CatalogSkidkiNatsenkiTsenovyeGruppysArgs) (*[]CatalogSkidkiNatsenkiTsenovyeGruppy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSkidkiNatsenkiTsenovyeGruppys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSkidkiNatsenkiTsenovyeGruppys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSkidkiNatsenkiTsenovyeGruppys(Where{})
}
func (r *GqlResolver) CatalogSkidkiNatsenkiNaborPodarkov(ctx context.Context, args CatalogSkidkiNatsenkiNaborPodarkovArgs) (*CatalogSkidkiNatsenkiNaborPodarkov, error) {
	return r.Client.CatalogSkidkiNatsenkiNaborPodarkov(args.Key, nil)
}
func (r *GqlResolver) CatalogSkidkiNatsenkiNaborPodarkovs(ctx context.Context, args CatalogSkidkiNatsenkiNaborPodarkovsArgs) (*[]CatalogSkidkiNatsenkiNaborPodarkov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSkidkiNatsenkiNaborPodarkovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSkidkiNatsenkiNaborPodarkovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSkidkiNatsenkiNaborPodarkovs(Where{})
}
func (r *GqlResolver) CatalogGruppyTsvetov(ctx context.Context, args CatalogGruppyTsvetovArgs) (*CatalogGruppyTsvetov, error) {
	return r.Client.CatalogGruppyTsvetov(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyTsvetovs(ctx context.Context, args CatalogGruppyTsvetovsArgs) (*[]CatalogGruppyTsvetov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyTsvetovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyTsvetovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyTsvetovs(Where{})
}
func (r *GqlResolver) DocumentDokumentRaschetovSKontragentom(ctx context.Context, args DocumentDokumentRaschetovSKontragentomArgs) (*DocumentDokumentRaschetovSKontragentom, error) {
	return r.Client.DocumentDokumentRaschetovSKontragentom(args.Key, nil)
}
func (r *GqlResolver) DocumentDokumentRaschetovSKontragentoms(ctx context.Context, args DocumentDokumentRaschetovSKontragentomsArgs) (*[]DocumentDokumentRaschetovSKontragentom, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentDokumentRaschetovSKontragentoms(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentDokumentRaschetovSKontragentoms(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentDokumentRaschetovSKontragentoms(Where{})
}
func (r *GqlResolver) CatalogDogovoryEkvairinga(ctx context.Context, args CatalogDogovoryEkvairingaArgs) (*CatalogDogovoryEkvairinga, error) {
	return r.Client.CatalogDogovoryEkvairinga(args.Key, nil)
}
func (r *GqlResolver) CatalogDogovoryEkvairingas(ctx context.Context, args CatalogDogovoryEkvairingasArgs) (*[]CatalogDogovoryEkvairinga, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDogovoryEkvairingas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDogovoryEkvairingas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDogovoryEkvairingas(Where{})
}
func (r *GqlResolver) CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie(ctx context.Context, args CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanieArgs) (*CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie, error) {
	return r.Client.CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie(args.Key, nil)
}
func (r *GqlResolver) CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanies(ctx context.Context, args CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivaniesArgs) (*[]CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanie, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanies(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanies(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDogovoryEkvairingaTarifyZaRaschetnoeObsluzhivanies(Where{})
}
func (r *GqlResolver) CatalogKachestvo(ctx context.Context, args CatalogKachestvoArgs) (*CatalogKachestvo, error) {
	return r.Client.CatalogKachestvo(args.Key, nil)
}
func (r *GqlResolver) CatalogKachestvos(ctx context.Context, args CatalogKachestvosArgs) (*[]CatalogKachestvo, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKachestvos(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKachestvos(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKachestvos(Where{})
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyKontragentov(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyKontragentovArgs) (*DocumentUstanovkaTsenNomenklaturyKontragentov, error) {
	return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentov(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyKontragentovs(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyKontragentovsArgs) (*[]DocumentUstanovkaTsenNomenklaturyKontragentov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovs(Where{})
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsenArgs) (*DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen, error) {
	return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsens(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsensArgs) (*[]DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsen, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsens(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsens(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovTipyTsens(Where{})
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyKontragentovTovary(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyKontragentovTovaryArgs) (*DocumentUstanovkaTsenNomenklaturyKontragentovTovary, error) {
	return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaTsenNomenklaturyKontragentovTovarys(ctx context.Context, args DocumentUstanovkaTsenNomenklaturyKontragentovTovarysArgs) (*[]DocumentUstanovkaTsenNomenklaturyKontragentovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaTsenNomenklaturyKontragentovTovarys(Where{})
}
func (r *GqlResolver) DocumentProtsentPoterPoDavaltsam(ctx context.Context, args DocumentProtsentPoterPoDavaltsamArgs) (*DocumentProtsentPoterPoDavaltsam, error) {
	return r.Client.DocumentProtsentPoterPoDavaltsam(args.Key, nil)
}
func (r *GqlResolver) DocumentProtsentPoterPoDavaltsams(ctx context.Context, args DocumentProtsentPoterPoDavaltsamsArgs) (*[]DocumentProtsentPoterPoDavaltsam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentProtsentPoterPoDavaltsams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentProtsentPoterPoDavaltsams(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentProtsentPoterPoDavaltsams(Where{})
}
func (r *GqlResolver) DocumentProtsentPoterPoDavaltsamProtsenty(ctx context.Context, args DocumentProtsentPoterPoDavaltsamProtsentyArgs) (*DocumentProtsentPoterPoDavaltsamProtsenty, error) {
	return r.Client.DocumentProtsentPoterPoDavaltsamProtsenty(args.Key, nil)
}
func (r *GqlResolver) DocumentProtsentPoterPoDavaltsamProtsentys(ctx context.Context, args DocumentProtsentPoterPoDavaltsamProtsentysArgs) (*[]DocumentProtsentPoterPoDavaltsamProtsenty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentProtsentPoterPoDavaltsamProtsentys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentProtsentPoterPoDavaltsamProtsentys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentProtsentPoterPoDavaltsamProtsentys(Where{})
}
func (r *GqlResolver) CatalogTovarnyePozitsii(ctx context.Context, args CatalogTovarnyePozitsiiArgs) (*CatalogTovarnyePozitsii, error) {
	return r.Client.CatalogTovarnyePozitsii(args.Key, nil)
}
func (r *GqlResolver) CatalogTovarnyePozitsiis(ctx context.Context, args CatalogTovarnyePozitsiisArgs) (*[]CatalogTovarnyePozitsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTovarnyePozitsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTovarnyePozitsiis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTovarnyePozitsiis(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieIskhodiashchee(ctx context.Context, args DocumentPlatezhnoePoruchenieIskhodiashcheeArgs) (*DocumentPlatezhnoePoruchenieIskhodiashchee, error) {
	return r.Client.DocumentPlatezhnoePoruchenieIskhodiashchee(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieIskhodiashchees(ctx context.Context, args DocumentPlatezhnoePoruchenieIskhodiashcheesArgs) (*[]DocumentPlatezhnoePoruchenieIskhodiashchee, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoePoruchenieIskhodiashchees(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoePoruchenieIskhodiashchees(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoePoruchenieIskhodiashchees(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha(ctx context.Context, args DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhaArgs) (*DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha, error) {
	return r.Client.DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhas(ctx context.Context, args DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhasArgs) (*[]DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoePoruchenieIskhodiashcheeRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta(ctx context.Context, args DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentaArgs) (*DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta, error) {
	return r.Client.DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta(args.Key, nil)
}
func (r *GqlResolver) DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentas(ctx context.Context, args DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentasArgs) (*[]DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlatezhnoePoruchenieIskhodiashcheeRekvizityKontragentas(Where{})
}
func (r *GqlResolver) CatalogfmOrganizatsionnoPravovyeFormy(ctx context.Context, args CatalogfmOrganizatsionnoPravovyeFormyArgs) (*CatalogfmOrganizatsionnoPravovyeFormy, error) {
	return r.Client.CatalogfmOrganizatsionnoPravovyeFormy(args.Key, nil)
}
func (r *GqlResolver) CatalogfmOrganizatsionnoPravovyeFormys(ctx context.Context, args CatalogfmOrganizatsionnoPravovyeFormysArgs) (*[]CatalogfmOrganizatsionnoPravovyeFormy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogfmOrganizatsionnoPravovyeFormys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogfmOrganizatsionnoPravovyeFormys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogfmOrganizatsionnoPravovyeFormys(Where{})
}
func (r *GqlResolver) CatalogTipyTsenNomenklatury(ctx context.Context, args CatalogTipyTsenNomenklaturyArgs) (*CatalogTipyTsenNomenklatury, error) {
	return r.Client.CatalogTipyTsenNomenklatury(args.Key, nil)
}
func (r *GqlResolver) CatalogTipyTsenNomenklaturys(ctx context.Context, args CatalogTipyTsenNomenklaturysArgs) (*[]CatalogTipyTsenNomenklatury, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTipyTsenNomenklaturys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTipyTsenNomenklaturys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTipyTsenNomenklaturys(Where{})
}
func (r *GqlResolver) CatalogStatiOtchetaPoProdazham(ctx context.Context, args CatalogStatiOtchetaPoProdazhamArgs) (*CatalogStatiOtchetaPoProdazham, error) {
	return r.Client.CatalogStatiOtchetaPoProdazham(args.Key, nil)
}
func (r *GqlResolver) CatalogStatiOtchetaPoProdazhams(ctx context.Context, args CatalogStatiOtchetaPoProdazhamsArgs) (*[]CatalogStatiOtchetaPoProdazham, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogStatiOtchetaPoProdazhams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogStatiOtchetaPoProdazhams(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogStatiOtchetaPoProdazhams(Where{})
}
func (r *GqlResolver) CatalogVidyKodirovokiIzdelii(ctx context.Context, args CatalogVidyKodirovokiIzdeliiArgs) (*CatalogVidyKodirovokiIzdelii, error) {
	return r.Client.CatalogVidyKodirovokiIzdelii(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyKodirovokiIzdeliis(ctx context.Context, args CatalogVidyKodirovokiIzdeliisArgs) (*[]CatalogVidyKodirovokiIzdelii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyKodirovokiIzdeliis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyKodirovokiIzdeliis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyKodirovokiIzdeliis(Where{})
}
func (r *GqlResolver) CatalogVidyKodirovokiIzdeliiElementyKodirovki(ctx context.Context, args CatalogVidyKodirovokiIzdeliiElementyKodirovkiArgs) (*CatalogVidyKodirovokiIzdeliiElementyKodirovki, error) {
	return r.Client.CatalogVidyKodirovokiIzdeliiElementyKodirovki(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyKodirovokiIzdeliiElementyKodirovkis(ctx context.Context, args CatalogVidyKodirovokiIzdeliiElementyKodirovkisArgs) (*[]CatalogVidyKodirovokiIzdeliiElementyKodirovki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyKodirovokiIzdeliiElementyKodirovkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyKodirovokiIzdeliiElementyKodirovkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyKodirovokiIzdeliiElementyKodirovkis(Where{})
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklatury(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturyArgs) (*DocumentUstanovkaSkidokNomenklatury, error) {
	return r.Client.DocumentUstanovkaSkidokNomenklatury(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklaturys(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturysArgs) (*[]DocumentUstanovkaSkidokNomenklatury, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaSkidokNomenklaturys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaSkidokNomenklaturys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaSkidokNomenklaturys(Where{})
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeliArgs) (*DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli, error) {
	return r.Client.DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedelis(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedelisArgs) (*[]DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedeli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedelis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedelis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaSkidokNomenklaturyVremiaPoDniamNedelis(Where{})
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartyArgs) (*DocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty, error) {
	return r.Client.DocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartys(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartysArgs) (*[]DocumentUstanovkaSkidokNomenklaturyDiskontnyeKarty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaSkidokNomenklaturyDiskontnyeKartys(Where{})
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkiArgs) (*DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki, error) {
	return r.Client.DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkis(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkisArgs) (*[]DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaSkidokNomenklaturyPoluchateliSkidkis(Where{})
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklaturyTovary(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturyTovaryArgs) (*DocumentUstanovkaSkidokNomenklaturyTovary, error) {
	return r.Client.DocumentUstanovkaSkidokNomenklaturyTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentUstanovkaSkidokNomenklaturyTovarys(ctx context.Context, args DocumentUstanovkaSkidokNomenklaturyTovarysArgs) (*[]DocumentUstanovkaSkidokNomenklaturyTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentUstanovkaSkidokNomenklaturyTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentUstanovkaSkidokNomenklaturyTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentUstanovkaSkidokNomenklaturyTovarys(Where{})
}
func (r *GqlResolver) CatalogUsloviiaPredostavleniiaSkidokNatsenok(ctx context.Context, args CatalogUsloviiaPredostavleniiaSkidokNatsenokArgs) (*CatalogUsloviiaPredostavleniiaSkidokNatsenok, error) {
	return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenok(args.Key, nil)
}
func (r *GqlResolver) CatalogUsloviiaPredostavleniiaSkidokNatsenoks(ctx context.Context, args CatalogUsloviiaPredostavleniiaSkidokNatsenoksArgs) (*[]CatalogUsloviiaPredostavleniiaSkidokNatsenok, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenoks(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenoks(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenoks(Where{})
}
func (r *GqlResolver) CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia(ctx context.Context, args CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviiaArgs) (*CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia, error) {
	return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia(args.Key, nil)
}
func (r *GqlResolver) CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviias(ctx context.Context, args CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviiasArgs) (*[]CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokVremiaDeistviias(Where{})
}
func (r *GqlResolver) CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli(ctx context.Context, args CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateliArgs) (*CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli, error) {
	return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli(args.Key, nil)
}
func (r *GqlResolver) CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchatelis(ctx context.Context, args CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchatelisArgs) (*[]CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchateli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchatelis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchatelis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokPoluchatelis(Where{})
}
func (r *GqlResolver) CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki(ctx context.Context, args CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkiArgs) (*CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki, error) {
	return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki(args.Key, nil)
}
func (r *GqlResolver) CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkis(ctx context.Context, args CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkisArgs) (*[]CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogUsloviiaPredostavleniiaSkidokNatsenokKomplektPokupkis(Where{})
}
func (r *GqlResolver) OutPay(ctx context.Context, args OutPayArgs) (*OutPay, error) {
	return r.Client.OutPay(args.Key, nil)
}
func (r *GqlResolver) OutPays(ctx context.Context, args OutPaysArgs) (*[]OutPay, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.OutPays(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.OutPays(Where{Filter: *args.Filter})
	}
	return r.Client.OutPays(Where{})
}
func (r *GqlResolver) DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha(ctx context.Context, args DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhaArgs) (*DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha, error) {
	return r.Client.DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha(args.Key, nil)
}
func (r *GqlResolver) DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhas(ctx context.Context, args DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhasArgs) (*[]DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezha, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRaskhodnyiKassovyiOrderRasshifrovkaPlatezhas(Where{})
}
func (r *GqlResolver) DocumentRaskhodnyiKassovyiOrderOplata(ctx context.Context, args DocumentRaskhodnyiKassovyiOrderOplataArgs) (*DocumentRaskhodnyiKassovyiOrderOplata, error) {
	return r.Client.DocumentRaskhodnyiKassovyiOrderOplata(args.Key, nil)
}
func (r *GqlResolver) DocumentRaskhodnyiKassovyiOrderOplatas(ctx context.Context, args DocumentRaskhodnyiKassovyiOrderOplatasArgs) (*[]DocumentRaskhodnyiKassovyiOrderOplata, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRaskhodnyiKassovyiOrderOplatas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRaskhodnyiKassovyiOrderOplatas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRaskhodnyiKassovyiOrderOplatas(Where{})
}
func (r *GqlResolver) DocumentRaskhodnyiKassovyiOrderTovary(ctx context.Context, args DocumentRaskhodnyiKassovyiOrderTovaryArgs) (*DocumentRaskhodnyiKassovyiOrderTovary, error) {
	return r.Client.DocumentRaskhodnyiKassovyiOrderTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentRaskhodnyiKassovyiOrderTovarys(ctx context.Context, args DocumentRaskhodnyiKassovyiOrderTovarysArgs) (*[]DocumentRaskhodnyiKassovyiOrderTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRaskhodnyiKassovyiOrderTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRaskhodnyiKassovyiOrderTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRaskhodnyiKassovyiOrderTovarys(Where{})
}
func (r *GqlResolver) DocumentSchetNaOplatuPostavshchika(ctx context.Context, args DocumentSchetNaOplatuPostavshchikaArgs) (*DocumentSchetNaOplatuPostavshchika, error) {
	return r.Client.DocumentSchetNaOplatuPostavshchika(args.Key, nil)
}
func (r *GqlResolver) DocumentSchetNaOplatuPostavshchikas(ctx context.Context, args DocumentSchetNaOplatuPostavshchikasArgs) (*[]DocumentSchetNaOplatuPostavshchika, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSchetNaOplatuPostavshchikas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSchetNaOplatuPostavshchikas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSchetNaOplatuPostavshchikas(Where{})
}
func (r *GqlResolver) DocumentSchetNaOplatuPostavshchikaTovary(ctx context.Context, args DocumentSchetNaOplatuPostavshchikaTovaryArgs) (*DocumentSchetNaOplatuPostavshchikaTovary, error) {
	return r.Client.DocumentSchetNaOplatuPostavshchikaTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentSchetNaOplatuPostavshchikaTovarys(ctx context.Context, args DocumentSchetNaOplatuPostavshchikaTovarysArgs) (*[]DocumentSchetNaOplatuPostavshchikaTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSchetNaOplatuPostavshchikaTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSchetNaOplatuPostavshchikaTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSchetNaOplatuPostavshchikaTovarys(Where{})
}
func (r *GqlResolver) DocumentSchetNaOplatuPostavshchikaUslugi(ctx context.Context, args DocumentSchetNaOplatuPostavshchikaUslugiArgs) (*DocumentSchetNaOplatuPostavshchikaUslugi, error) {
	return r.Client.DocumentSchetNaOplatuPostavshchikaUslugi(args.Key, nil)
}
func (r *GqlResolver) DocumentSchetNaOplatuPostavshchikaUslugis(ctx context.Context, args DocumentSchetNaOplatuPostavshchikaUslugisArgs) (*[]DocumentSchetNaOplatuPostavshchikaUslugi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSchetNaOplatuPostavshchikaUslugis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSchetNaOplatuPostavshchikaUslugis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSchetNaOplatuPostavshchikaUslugis(Where{})
}
func (r *GqlResolver) DocumentReestrSpetssviaz(ctx context.Context, args DocumentReestrSpetssviazArgs) (*DocumentReestrSpetssviaz, error) {
	return r.Client.DocumentReestrSpetssviaz(args.Key, nil)
}
func (r *GqlResolver) DocumentReestrSpetssviazs(ctx context.Context, args DocumentReestrSpetssviazsArgs) (*[]DocumentReestrSpetssviaz, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentReestrSpetssviazs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentReestrSpetssviazs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentReestrSpetssviazs(Where{})
}
func (r *GqlResolver) DocumentReestrSpetssviazKlienty(ctx context.Context, args DocumentReestrSpetssviazKlientyArgs) (*DocumentReestrSpetssviazKlienty, error) {
	return r.Client.DocumentReestrSpetssviazKlienty(args.Key, nil)
}
func (r *GqlResolver) DocumentReestrSpetssviazKlientys(ctx context.Context, args DocumentReestrSpetssviazKlientysArgs) (*[]DocumentReestrSpetssviazKlienty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentReestrSpetssviazKlientys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentReestrSpetssviazKlientys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentReestrSpetssviazKlientys(Where{})
}
func (r *GqlResolver) DocumentJournalKassovyeDokumenty(ctx context.Context, args DocumentJournalKassovyeDokumentyArgs) (*DocumentJournalKassovyeDokumenty, error) {
	return r.Client.DocumentJournalKassovyeDokumenty(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalKassovyeDokumentys(ctx context.Context, args DocumentJournalKassovyeDokumentysArgs) (*[]DocumentJournalKassovyeDokumenty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalKassovyeDokumentys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalKassovyeDokumentys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalKassovyeDokumentys(Where{})
}
func (r *GqlResolver) InitialInstance(ctx context.Context, args InitialInstanceArgs) (*InitialInstance, error) {
	return r.Client.InitialInstance(args.Key, nil)
}
func (r *GqlResolver) InitialInstances(ctx context.Context, args InitialInstancesArgs) (*[]InitialInstance, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.InitialInstances(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.InitialInstances(Where{Filter: *args.Filter})
	}
	return r.Client.InitialInstances(Where{})
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovVzaimoraschety(ctx context.Context, args DocumentVvodNachalnykhOstatkovVzaimoraschetyArgs) (*DocumentVvodNachalnykhOstatkovVzaimoraschety, error) {
	return r.Client.DocumentVvodNachalnykhOstatkovVzaimoraschety(args.Key, nil)
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovVzaimoraschetys(ctx context.Context, args DocumentVvodNachalnykhOstatkovVzaimoraschetysArgs) (*[]DocumentVvodNachalnykhOstatkovVzaimoraschety, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVvodNachalnykhOstatkovVzaimoraschetys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVvodNachalnykhOstatkovVzaimoraschetys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVvodNachalnykhOstatkovVzaimoraschetys(Where{})
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovTovary(ctx context.Context, args DocumentVvodNachalnykhOstatkovTovaryArgs) (*DocumentVvodNachalnykhOstatkovTovary, error) {
	return r.Client.DocumentVvodNachalnykhOstatkovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentVvodNachalnykhOstatkovTovarys(ctx context.Context, args DocumentVvodNachalnykhOstatkovTovarysArgs) (*[]DocumentVvodNachalnykhOstatkovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVvodNachalnykhOstatkovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVvodNachalnykhOstatkovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVvodNachalnykhOstatkovTovarys(Where{})
}
func (r *GqlResolver) Posting(ctx context.Context, args PostingArgs) (*Posting, error) {
	return r.Client.Posting(args.Key, nil)
}
func (r *GqlResolver) Postings(ctx context.Context, args PostingsArgs) (*[]Posting, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Postings(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Postings(Where{Filter: *args.Filter})
	}
	return r.Client.Postings(Where{})
}
func (r *GqlResolver) DocumentOprikhodovanieTovarovTovary(ctx context.Context, args DocumentOprikhodovanieTovarovTovaryArgs) (*DocumentOprikhodovanieTovarovTovary, error) {
	return r.Client.DocumentOprikhodovanieTovarovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentOprikhodovanieTovarovTovarys(ctx context.Context, args DocumentOprikhodovanieTovarovTovarysArgs) (*[]DocumentOprikhodovanieTovarovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOprikhodovanieTovarovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOprikhodovanieTovarovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOprikhodovanieTovarovTovarys(Where{})
}
func (r *GqlResolver) DocumentOprikhodovanieTovarovSertifikaty(ctx context.Context, args DocumentOprikhodovanieTovarovSertifikatyArgs) (*DocumentOprikhodovanieTovarovSertifikaty, error) {
	return r.Client.DocumentOprikhodovanieTovarovSertifikaty(args.Key, nil)
}
func (r *GqlResolver) DocumentOprikhodovanieTovarovSertifikatys(ctx context.Context, args DocumentOprikhodovanieTovarovSertifikatysArgs) (*[]DocumentOprikhodovanieTovarovSertifikaty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOprikhodovanieTovarovSertifikatys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOprikhodovanieTovarovSertifikatys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOprikhodovanieTovarovSertifikatys(Where{})
}
func (r *GqlResolver) CatalogKomplekty(ctx context.Context, args CatalogKomplektyArgs) (*CatalogKomplekty, error) {
	return r.Client.CatalogKomplekty(args.Key, nil)
}
func (r *GqlResolver) CatalogKomplektys(ctx context.Context, args CatalogKomplektysArgs) (*[]CatalogKomplekty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKomplektys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKomplektys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKomplektys(Where{})
}
func (r *GqlResolver) DocumentPereotsenkaTovarovPriniatykhNaKomissiiu(ctx context.Context, args DocumentPereotsenkaTovarovPriniatykhNaKomissiiuArgs) (*DocumentPereotsenkaTovarovPriniatykhNaKomissiiu, error) {
	return r.Client.DocumentPereotsenkaTovarovPriniatykhNaKomissiiu(args.Key, nil)
}
func (r *GqlResolver) DocumentPereotsenkaTovarovPriniatykhNaKomissiius(ctx context.Context, args DocumentPereotsenkaTovarovPriniatykhNaKomissiiusArgs) (*[]DocumentPereotsenkaTovarovPriniatykhNaKomissiiu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPereotsenkaTovarovPriniatykhNaKomissiius(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPereotsenkaTovarovPriniatykhNaKomissiius(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPereotsenkaTovarovPriniatykhNaKomissiius(Where{})
}
func (r *GqlResolver) DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary(ctx context.Context, args DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovaryArgs) (*DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary, error) {
	return r.Client.DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovarys(ctx context.Context, args DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovarysArgs) (*[]DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPereotsenkaTovarovPriniatykhNaKomissiiuTovarys(Where{})
}
func (r *GqlResolver) DocumentElektronnoePismo(ctx context.Context, args DocumentElektronnoePismoArgs) (*DocumentElektronnoePismo, error) {
	return r.Client.DocumentElektronnoePismo(args.Key, nil)
}
func (r *GqlResolver) DocumentElektronnoePismos(ctx context.Context, args DocumentElektronnoePismosArgs) (*[]DocumentElektronnoePismo, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentElektronnoePismos(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentElektronnoePismos(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentElektronnoePismos(Where{})
}
func (r *GqlResolver) DocumentElektronnoePismoKomuTCh(ctx context.Context, args DocumentElektronnoePismoKomuTChArgs) (*DocumentElektronnoePismoKomuTCh, error) {
	return r.Client.DocumentElektronnoePismoKomuTCh(args.Key, nil)
}
func (r *GqlResolver) DocumentElektronnoePismoKomuTChs(ctx context.Context, args DocumentElektronnoePismoKomuTChsArgs) (*[]DocumentElektronnoePismoKomuTCh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentElektronnoePismoKomuTChs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentElektronnoePismoKomuTChs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentElektronnoePismoKomuTChs(Where{})
}
func (r *GqlResolver) DocumentElektronnoePismoKopiiTCh(ctx context.Context, args DocumentElektronnoePismoKopiiTChArgs) (*DocumentElektronnoePismoKopiiTCh, error) {
	return r.Client.DocumentElektronnoePismoKopiiTCh(args.Key, nil)
}
func (r *GqlResolver) DocumentElektronnoePismoKopiiTChs(ctx context.Context, args DocumentElektronnoePismoKopiiTChsArgs) (*[]DocumentElektronnoePismoKopiiTCh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentElektronnoePismoKopiiTChs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentElektronnoePismoKopiiTChs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentElektronnoePismoKopiiTChs(Where{})
}
func (r *GqlResolver) DocumentElektronnoePismoSkrytyeKopiiTCh(ctx context.Context, args DocumentElektronnoePismoSkrytyeKopiiTChArgs) (*DocumentElektronnoePismoSkrytyeKopiiTCh, error) {
	return r.Client.DocumentElektronnoePismoSkrytyeKopiiTCh(args.Key, nil)
}
func (r *GqlResolver) DocumentElektronnoePismoSkrytyeKopiiTChs(ctx context.Context, args DocumentElektronnoePismoSkrytyeKopiiTChsArgs) (*[]DocumentElektronnoePismoSkrytyeKopiiTCh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentElektronnoePismoSkrytyeKopiiTChs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentElektronnoePismoSkrytyeKopiiTChs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentElektronnoePismoSkrytyeKopiiTChs(Where{})
}
func (r *GqlResolver) CatalogGruppyDefektov(ctx context.Context, args CatalogGruppyDefektovArgs) (*CatalogGruppyDefektov, error) {
	return r.Client.CatalogGruppyDefektov(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyDefektovs(ctx context.Context, args CatalogGruppyDefektovsArgs) (*[]CatalogGruppyDefektov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyDefektovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyDefektovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyDefektovs(Where{})
}
func (r *GqlResolver) CatalogfmAnketaKlientaBenefitsariia(ctx context.Context, args CatalogfmAnketaKlientaBenefitsariiaArgs) (*CatalogfmAnketaKlientaBenefitsariia, error) {
	return r.Client.CatalogfmAnketaKlientaBenefitsariia(args.Key, nil)
}
func (r *GqlResolver) CatalogfmAnketaKlientaBenefitsariias(ctx context.Context, args CatalogfmAnketaKlientaBenefitsariiasArgs) (*[]CatalogfmAnketaKlientaBenefitsariia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogfmAnketaKlientaBenefitsariias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogfmAnketaKlientaBenefitsariias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogfmAnketaKlientaBenefitsariias(Where{})
}
func (r *GqlResolver) CatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta(ctx context.Context, args CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentaArgs) (*CatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta, error) {
	return r.Client.CatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta(args.Key, nil)
}
func (r *GqlResolver) CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentas(ctx context.Context, args CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentasArgs) (*[]CatalogfmAnketaKlientaBenefitsariiaDannyeKontragenta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogfmAnketaKlientaBenefitsariiaDannyeKontragentas(Where{})
}
func (r *GqlResolver) CatalogTsenovyeGruppy(ctx context.Context, args CatalogTsenovyeGruppyArgs) (*CatalogTsenovyeGruppy, error) {
	return r.Client.CatalogTsenovyeGruppy(args.Key, nil)
}
func (r *GqlResolver) CatalogTsenovyeGruppys(ctx context.Context, args CatalogTsenovyeGruppysArgs) (*[]CatalogTsenovyeGruppy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTsenovyeGruppys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTsenovyeGruppys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTsenovyeGruppys(Where{})
}
func (r *GqlResolver) CatalogPravilaTsenoobrazovaniia(ctx context.Context, args CatalogPravilaTsenoobrazovaniiaArgs) (*CatalogPravilaTsenoobrazovaniia, error) {
	return r.Client.CatalogPravilaTsenoobrazovaniia(args.Key, nil)
}
func (r *GqlResolver) CatalogPravilaTsenoobrazovaniias(ctx context.Context, args CatalogPravilaTsenoobrazovaniiasArgs) (*[]CatalogPravilaTsenoobrazovaniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogPravilaTsenoobrazovaniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogPravilaTsenoobrazovaniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogPravilaTsenoobrazovaniias(Where{})
}
func (r *GqlResolver) CatalogPravilaTsenoobrazovaniiaTsenovyeGruppy(ctx context.Context, args CatalogPravilaTsenoobrazovaniiaTsenovyeGruppyArgs) (*CatalogPravilaTsenoobrazovaniiaTsenovyeGruppy, error) {
	return r.Client.CatalogPravilaTsenoobrazovaniiaTsenovyeGruppy(args.Key, nil)
}
func (r *GqlResolver) CatalogPravilaTsenoobrazovaniiaTsenovyeGruppys(ctx context.Context, args CatalogPravilaTsenoobrazovaniiaTsenovyeGruppysArgs) (*[]CatalogPravilaTsenoobrazovaniiaTsenovyeGruppy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogPravilaTsenoobrazovaniiaTsenovyeGruppys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogPravilaTsenoobrazovaniiaTsenovyeGruppys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogPravilaTsenoobrazovaniiaTsenovyeGruppys(Where{})
}
func (r *GqlResolver) DocumentObieiavlenieNaVznosNalichnymi(ctx context.Context, args DocumentObieiavlenieNaVznosNalichnymiArgs) (*DocumentObieiavlenieNaVznosNalichnymi, error) {
	return r.Client.DocumentObieiavlenieNaVznosNalichnymi(args.Key, nil)
}
func (r *GqlResolver) DocumentObieiavlenieNaVznosNalichnymis(ctx context.Context, args DocumentObieiavlenieNaVznosNalichnymisArgs) (*[]DocumentObieiavlenieNaVznosNalichnymi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentObieiavlenieNaVznosNalichnymis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentObieiavlenieNaVznosNalichnymis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentObieiavlenieNaVznosNalichnymis(Where{})
}
func (r *GqlResolver) CatalogValiuty(ctx context.Context, args CatalogValiutyArgs) (*CatalogValiuty, error) {
	return r.Client.CatalogValiuty(args.Key, nil)
}
func (r *GqlResolver) CatalogValiutys(ctx context.Context, args CatalogValiutysArgs) (*[]CatalogValiuty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogValiutys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogValiutys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogValiutys(Where{})
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku(ctx context.Context, args DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuArgs) (*DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku, error) {
	return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku(args.Key, nil)
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkus(ctx context.Context, args DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkusArgs) (*[]DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochku, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkus(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkus(Where{})
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary(ctx context.Context, args DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovaryArgs) (*DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary, error) {
	return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovarys(ctx context.Context, args DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovarysArgs) (*[]DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuTovarys(Where{})
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi(ctx context.Context, args DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugiArgs) (*DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi, error) {
	return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi(args.Key, nil)
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugis(ctx context.Context, args DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugisArgs) (*[]DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPostuplenieTovarovUslugVNeavtomatizirovannuiuTorgovuiuTochkuUslugis(Where{})
}
func (r *GqlResolver) CatalogKassyKKM(ctx context.Context, args CatalogKassyKKMArgs) (*CatalogKassyKKM, error) {
	return r.Client.CatalogKassyKKM(args.Key, nil)
}
func (r *GqlResolver) CatalogKassyKKMs(ctx context.Context, args CatalogKassyKKMsArgs) (*[]CatalogKassyKKM, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKassyKKMs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKassyKKMs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKassyKKMs(Where{})
}
func (r *GqlResolver) Probe(ctx context.Context, args ProbeArgs) (*Probe, error) {
	return r.Client.Probe(args.Key, nil)
}
func (r *GqlResolver) Probes(ctx context.Context, args ProbesArgs) (*[]Probe, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Probes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Probes(Where{Filter: *args.Filter})
	}
	return r.Client.Probes(Where{})
}
func (r *GqlResolver) CatalogGruppyDostupa(ctx context.Context, args CatalogGruppyDostupaArgs) (*CatalogGruppyDostupa, error) {
	return r.Client.CatalogGruppyDostupa(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyDostupas(ctx context.Context, args CatalogGruppyDostupasArgs) (*[]CatalogGruppyDostupa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyDostupas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyDostupas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyDostupas(Where{})
}
func (r *GqlResolver) CatalogGruppyDostupaPolzovateli(ctx context.Context, args CatalogGruppyDostupaPolzovateliArgs) (*CatalogGruppyDostupaPolzovateli, error) {
	return r.Client.CatalogGruppyDostupaPolzovateli(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyDostupaPolzovatelis(ctx context.Context, args CatalogGruppyDostupaPolzovatelisArgs) (*[]CatalogGruppyDostupaPolzovateli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyDostupaPolzovatelis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyDostupaPolzovatelis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyDostupaPolzovatelis(Where{})
}
func (r *GqlResolver) CatalogGruppyDostupaVidyDostupa(ctx context.Context, args CatalogGruppyDostupaVidyDostupaArgs) (*CatalogGruppyDostupaVidyDostupa, error) {
	return r.Client.CatalogGruppyDostupaVidyDostupa(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyDostupaVidyDostupas(ctx context.Context, args CatalogGruppyDostupaVidyDostupasArgs) (*[]CatalogGruppyDostupaVidyDostupa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyDostupaVidyDostupas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyDostupaVidyDostupas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyDostupaVidyDostupas(Where{})
}
func (r *GqlResolver) CatalogGruppyDostupaZnacheniiaDostupa(ctx context.Context, args CatalogGruppyDostupaZnacheniiaDostupaArgs) (*CatalogGruppyDostupaZnacheniiaDostupa, error) {
	return r.Client.CatalogGruppyDostupaZnacheniiaDostupa(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyDostupaZnacheniiaDostupas(ctx context.Context, args CatalogGruppyDostupaZnacheniiaDostupasArgs) (*[]CatalogGruppyDostupaZnacheniiaDostupa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyDostupaZnacheniiaDostupas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyDostupaZnacheniiaDostupas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyDostupaZnacheniiaDostupas(Where{})
}
func (r *GqlResolver) CatalogGruppyDostupaDostupPoPodsistemam(ctx context.Context, args CatalogGruppyDostupaDostupPoPodsistemamArgs) (*CatalogGruppyDostupaDostupPoPodsistemam, error) {
	return r.Client.CatalogGruppyDostupaDostupPoPodsistemam(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppyDostupaDostupPoPodsistemams(ctx context.Context, args CatalogGruppyDostupaDostupPoPodsistemamsArgs) (*[]CatalogGruppyDostupaDostupPoPodsistemam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppyDostupaDostupPoPodsistemams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppyDostupaDostupPoPodsistemams(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppyDostupaDostupPoPodsistemams(Where{})
}
func (r *GqlResolver) CatalogVidyKontaktnoiInformatsii(ctx context.Context, args CatalogVidyKontaktnoiInformatsiiArgs) (*CatalogVidyKontaktnoiInformatsii, error) {
	return r.Client.CatalogVidyKontaktnoiInformatsii(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyKontaktnoiInformatsiis(ctx context.Context, args CatalogVidyKontaktnoiInformatsiisArgs) (*[]CatalogVidyKontaktnoiInformatsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyKontaktnoiInformatsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyKontaktnoiInformatsiis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyKontaktnoiInformatsiis(Where{})
}
func (r *GqlResolver) CatalogNomenklaturnyeGruppy(ctx context.Context, args CatalogNomenklaturnyeGruppyArgs) (*CatalogNomenklaturnyeGruppy, error) {
	return r.Client.CatalogNomenklaturnyeGruppy(args.Key, nil)
}
func (r *GqlResolver) CatalogNomenklaturnyeGruppys(ctx context.Context, args CatalogNomenklaturnyeGruppysArgs) (*[]CatalogNomenklaturnyeGruppy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNomenklaturnyeGruppys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNomenklaturnyeGruppys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNomenklaturnyeGruppys(Where{})
}
func (r *GqlResolver) DocumentReestrSchetov(ctx context.Context, args DocumentReestrSchetovArgs) (*DocumentReestrSchetov, error) {
	return r.Client.DocumentReestrSchetov(args.Key, nil)
}
func (r *GqlResolver) DocumentReestrSchetovs(ctx context.Context, args DocumentReestrSchetovsArgs) (*[]DocumentReestrSchetov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentReestrSchetovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentReestrSchetovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentReestrSchetovs(Where{})
}
func (r *GqlResolver) DocumentReestrSchetovReestr(ctx context.Context, args DocumentReestrSchetovReestrArgs) (*DocumentReestrSchetovReestr, error) {
	return r.Client.DocumentReestrSchetovReestr(args.Key, nil)
}
func (r *GqlResolver) DocumentReestrSchetovReestrs(ctx context.Context, args DocumentReestrSchetovReestrsArgs) (*[]DocumentReestrSchetovReestr, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentReestrSchetovReestrs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentReestrSchetovReestrs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentReestrSchetovReestrs(Where{})
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu(ctx context.Context, args DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuArgs) (*DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu, error) {
	return r.Client.DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu(args.Key, nil)
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovOtdannykhNaKomissiius(ctx context.Context, args DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiusArgs) (*[]DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiu, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInventarizatsiiaTovarovOtdannykhNaKomissiius(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInventarizatsiiaTovarovOtdannykhNaKomissiius(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInventarizatsiiaTovarovOtdannykhNaKomissiius(Where{})
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary(ctx context.Context, args DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovaryArgs) (*DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary, error) {
	return r.Client.DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovarys(ctx context.Context, args DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovarysArgs) (*[]DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentInventarizatsiiaTovarovOtdannykhNaKomissiiuTovarys(Where{})
}
func (r *GqlResolver) CatalogKlassifikatorStranMira(ctx context.Context, args CatalogKlassifikatorStranMiraArgs) (*CatalogKlassifikatorStranMira, error) {
	return r.Client.CatalogKlassifikatorStranMira(args.Key, nil)
}
func (r *GqlResolver) CatalogKlassifikatorStranMiras(ctx context.Context, args CatalogKlassifikatorStranMirasArgs) (*[]CatalogKlassifikatorStranMira, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKlassifikatorStranMiras(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKlassifikatorStranMiras(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKlassifikatorStranMiras(Where{})
}
func (r *GqlResolver) CatalogKlassifikatorEdinitsIzmereniia(ctx context.Context, args CatalogKlassifikatorEdinitsIzmereniiaArgs) (*CatalogKlassifikatorEdinitsIzmereniia, error) {
	return r.Client.CatalogKlassifikatorEdinitsIzmereniia(args.Key, nil)
}
func (r *GqlResolver) CatalogKlassifikatorEdinitsIzmereniias(ctx context.Context, args CatalogKlassifikatorEdinitsIzmereniiasArgs) (*[]CatalogKlassifikatorEdinitsIzmereniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKlassifikatorEdinitsIzmereniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKlassifikatorEdinitsIzmereniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKlassifikatorEdinitsIzmereniias(Where{})
}
func (r *GqlResolver) CatalogNastroikiRMK(ctx context.Context, args CatalogNastroikiRMKArgs) (*CatalogNastroikiRMK, error) {
	return r.Client.CatalogNastroikiRMK(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiRMKs(ctx context.Context, args CatalogNastroikiRMKsArgs) (*[]CatalogNastroikiRMK, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiRMKs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiRMKs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiRMKs(Where{})
}
func (r *GqlResolver) CatalogNastroikiRMKPoriadokPrimeneniiaSkidok(ctx context.Context, args CatalogNastroikiRMKPoriadokPrimeneniiaSkidokArgs) (*CatalogNastroikiRMKPoriadokPrimeneniiaSkidok, error) {
	return r.Client.CatalogNastroikiRMKPoriadokPrimeneniiaSkidok(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiRMKPoriadokPrimeneniiaSkidoks(ctx context.Context, args CatalogNastroikiRMKPoriadokPrimeneniiaSkidoksArgs) (*[]CatalogNastroikiRMKPoriadokPrimeneniiaSkidok, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiRMKPoriadokPrimeneniiaSkidoks(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiRMKPoriadokPrimeneniiaSkidoks(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiRMKPoriadokPrimeneniiaSkidoks(Where{})
}
func (r *GqlResolver) CatalogNastroikiRMKSostavNaimenovaniia(ctx context.Context, args CatalogNastroikiRMKSostavNaimenovaniiaArgs) (*CatalogNastroikiRMKSostavNaimenovaniia, error) {
	return r.Client.CatalogNastroikiRMKSostavNaimenovaniia(args.Key, nil)
}
func (r *GqlResolver) CatalogNastroikiRMKSostavNaimenovaniias(ctx context.Context, args CatalogNastroikiRMKSostavNaimenovaniiasArgs) (*[]CatalogNastroikiRMKSostavNaimenovaniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogNastroikiRMKSostavNaimenovaniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogNastroikiRMKSostavNaimenovaniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogNastroikiRMKSostavNaimenovaniias(Where{})
}
func (r *GqlResolver) CatalogKharakteristikiNomenklatury(ctx context.Context, args CatalogKharakteristikiNomenklaturyArgs) (*CatalogKharakteristikiNomenklatury, error) {
	return r.Client.CatalogKharakteristikiNomenklatury(args.Key, nil)
}
func (r *GqlResolver) CatalogKharakteristikiNomenklaturys(ctx context.Context, args CatalogKharakteristikiNomenklaturysArgs) (*[]CatalogKharakteristikiNomenklatury, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKharakteristikiNomenklaturys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKharakteristikiNomenklaturys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKharakteristikiNomenklaturys(Where{})
}
func (r *GqlResolver) CatalogKharakteristikiNomenklaturySpetsifikatsiia(ctx context.Context, args CatalogKharakteristikiNomenklaturySpetsifikatsiiaArgs) (*CatalogKharakteristikiNomenklaturySpetsifikatsiia, error) {
	return r.Client.CatalogKharakteristikiNomenklaturySpetsifikatsiia(args.Key, nil)
}
func (r *GqlResolver) CatalogKharakteristikiNomenklaturySpetsifikatsiias(ctx context.Context, args CatalogKharakteristikiNomenklaturySpetsifikatsiiasArgs) (*[]CatalogKharakteristikiNomenklaturySpetsifikatsiia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogKharakteristikiNomenklaturySpetsifikatsiias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogKharakteristikiNomenklaturySpetsifikatsiias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogKharakteristikiNomenklaturySpetsifikatsiias(Where{})
}
func (r *GqlResolver) DocumentOtborTovarov(ctx context.Context, args DocumentOtborTovarovArgs) (*DocumentOtborTovarov, error) {
	return r.Client.DocumentOtborTovarov(args.Key, nil)
}
func (r *GqlResolver) DocumentOtborTovarovs(ctx context.Context, args DocumentOtborTovarovsArgs) (*[]DocumentOtborTovarov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtborTovarovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtborTovarovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtborTovarovs(Where{})
}
func (r *GqlResolver) DocumentOtborTovarovTovary(ctx context.Context, args DocumentOtborTovarovTovaryArgs) (*DocumentOtborTovarovTovary, error) {
	return r.Client.DocumentOtborTovarovTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentOtborTovarovTovarys(ctx context.Context, args DocumentOtborTovarovTovarysArgs) (*[]DocumentOtborTovarovTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtborTovarovTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtborTovarovTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtborTovarovTovarys(Where{})
}
func (r *GqlResolver) DocumentOtborTovarovTovaryKOtboru(ctx context.Context, args DocumentOtborTovarovTovaryKOtboruArgs) (*DocumentOtborTovarovTovaryKOtboru, error) {
	return r.Client.DocumentOtborTovarovTovaryKOtboru(args.Key, nil)
}
func (r *GqlResolver) DocumentOtborTovarovTovaryKOtborus(ctx context.Context, args DocumentOtborTovarovTovaryKOtborusArgs) (*[]DocumentOtborTovarovTovaryKOtboru, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtborTovarovTovaryKOtborus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtborTovarovTovaryKOtborus(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtborTovarovTovaryKOtborus(Where{})
}
func (r *GqlResolver) CatalogSposobyDostavkiTovara(ctx context.Context, args CatalogSposobyDostavkiTovaraArgs) (*CatalogSposobyDostavkiTovara, error) {
	return r.Client.CatalogSposobyDostavkiTovara(args.Key, nil)
}
func (r *GqlResolver) CatalogSposobyDostavkiTovaras(ctx context.Context, args CatalogSposobyDostavkiTovarasArgs) (*[]CatalogSposobyDostavkiTovara, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSposobyDostavkiTovaras(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSposobyDostavkiTovaras(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSposobyDostavkiTovaras(Where{})
}
func (r *GqlResolver) CatalogPodrazdeleniia(ctx context.Context, args CatalogPodrazdeleniiaArgs) (*CatalogPodrazdeleniia, error) {
	return r.Client.CatalogPodrazdeleniia(args.Key, nil)
}
func (r *GqlResolver) CatalogPodrazdeleniias(ctx context.Context, args CatalogPodrazdeleniiasArgs) (*[]CatalogPodrazdeleniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogPodrazdeleniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogPodrazdeleniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogPodrazdeleniias(Where{})
}
func (r *GqlResolver) DocumentJournalPreiskuranty(ctx context.Context, args DocumentJournalPreiskurantyArgs) (*DocumentJournalPreiskuranty, error) {
	return r.Client.DocumentJournalPreiskuranty(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalPreiskurantys(ctx context.Context, args DocumentJournalPreiskurantysArgs) (*[]DocumentJournalPreiskuranty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalPreiskurantys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalPreiskurantys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalPreiskurantys(Where{})
}
func (r *GqlResolver) CatalogRelizyIuvelirnykhSalonov(ctx context.Context, args CatalogRelizyIuvelirnykhSalonovArgs) (*CatalogRelizyIuvelirnykhSalonov, error) {
	return r.Client.CatalogRelizyIuvelirnykhSalonov(args.Key, nil)
}
func (r *GqlResolver) CatalogRelizyIuvelirnykhSalonovs(ctx context.Context, args CatalogRelizyIuvelirnykhSalonovsArgs) (*[]CatalogRelizyIuvelirnykhSalonov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogRelizyIuvelirnykhSalonovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogRelizyIuvelirnykhSalonovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogRelizyIuvelirnykhSalonovs(Where{})
}
func (r *GqlResolver) CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy(ctx context.Context, args CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizyArgs) (*CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy, error) {
	return r.Client.CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy(args.Key, nil)
}
func (r *GqlResolver) CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizys(ctx context.Context, args CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizysArgs) (*[]CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogRelizyIuvelirnykhSalonovObnovliaemyeRelizys(Where{})
}
func (r *GqlResolver) DocumentOtchetKomissioneraOProdazhakh(ctx context.Context, args DocumentOtchetKomissioneraOProdazhakhArgs) (*DocumentOtchetKomissioneraOProdazhakh, error) {
	return r.Client.DocumentOtchetKomissioneraOProdazhakh(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetKomissioneraOProdazhakhs(ctx context.Context, args DocumentOtchetKomissioneraOProdazhakhsArgs) (*[]DocumentOtchetKomissioneraOProdazhakh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetKomissioneraOProdazhakhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetKomissioneraOProdazhakhs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetKomissioneraOProdazhakhs(Where{})
}
func (r *GqlResolver) DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva(ctx context.Context, args DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvaArgs) (*DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva, error) {
	return r.Client.DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvas(ctx context.Context, args DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvasArgs) (*[]DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstva, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetKomissioneraOProdazhakhDenezhnyeSredstvas(Where{})
}
func (r *GqlResolver) DocumentOtchetKomissioneraOProdazhakhTovary(ctx context.Context, args DocumentOtchetKomissioneraOProdazhakhTovaryArgs) (*DocumentOtchetKomissioneraOProdazhakhTovary, error) {
	return r.Client.DocumentOtchetKomissioneraOProdazhakhTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentOtchetKomissioneraOProdazhakhTovarys(ctx context.Context, args DocumentOtchetKomissioneraOProdazhakhTovarysArgs) (*[]DocumentOtchetKomissioneraOProdazhakhTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentOtchetKomissioneraOProdazhakhTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentOtchetKomissioneraOProdazhakhTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentOtchetKomissioneraOProdazhakhTovarys(Where{})
}
func (r *GqlResolver) CatalogTovarnyeKategorii(ctx context.Context, args CatalogTovarnyeKategoriiArgs) (*CatalogTovarnyeKategorii, error) {
	return r.Client.CatalogTovarnyeKategorii(args.Key, nil)
}
func (r *GqlResolver) CatalogTovarnyeKategoriis(ctx context.Context, args CatalogTovarnyeKategoriisArgs) (*[]CatalogTovarnyeKategorii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTovarnyeKategoriis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTovarnyeKategoriis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTovarnyeKategoriis(Where{})
}
func (r *GqlResolver) CatalogDokumentyUdostoveriaiushchieLichnost(ctx context.Context, args CatalogDokumentyUdostoveriaiushchieLichnostArgs) (*CatalogDokumentyUdostoveriaiushchieLichnost, error) {
	return r.Client.CatalogDokumentyUdostoveriaiushchieLichnost(args.Key, nil)
}
func (r *GqlResolver) CatalogDokumentyUdostoveriaiushchieLichnosts(ctx context.Context, args CatalogDokumentyUdostoveriaiushchieLichnostsArgs) (*[]CatalogDokumentyUdostoveriaiushchieLichnost, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogDokumentyUdostoveriaiushchieLichnosts(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogDokumentyUdostoveriaiushchieLichnosts(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogDokumentyUdostoveriaiushchieLichnosts(Where{})
}
func (r *GqlResolver) CatalogFiltryDliaElektronnykhPisem(ctx context.Context, args CatalogFiltryDliaElektronnykhPisemArgs) (*CatalogFiltryDliaElektronnykhPisem, error) {
	return r.Client.CatalogFiltryDliaElektronnykhPisem(args.Key, nil)
}
func (r *GqlResolver) CatalogFiltryDliaElektronnykhPisems(ctx context.Context, args CatalogFiltryDliaElektronnykhPisemsArgs) (*[]CatalogFiltryDliaElektronnykhPisem, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFiltryDliaElektronnykhPisems(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFiltryDliaElektronnykhPisems(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFiltryDliaElektronnykhPisems(Where{})
}
func (r *GqlResolver) CatalogFiltryDliaElektronnykhPisemDeistviiaFiltra(ctx context.Context, args CatalogFiltryDliaElektronnykhPisemDeistviiaFiltraArgs) (*CatalogFiltryDliaElektronnykhPisemDeistviiaFiltra, error) {
	return r.Client.CatalogFiltryDliaElektronnykhPisemDeistviiaFiltra(args.Key, nil)
}
func (r *GqlResolver) CatalogFiltryDliaElektronnykhPisemDeistviiaFiltras(ctx context.Context, args CatalogFiltryDliaElektronnykhPisemDeistviiaFiltrasArgs) (*[]CatalogFiltryDliaElektronnykhPisemDeistviiaFiltra, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFiltryDliaElektronnykhPisemDeistviiaFiltras(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFiltryDliaElektronnykhPisemDeistviiaFiltras(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFiltryDliaElektronnykhPisemDeistviiaFiltras(Where{})
}
func (r *GqlResolver) CatalogFiltryDliaElektronnykhPisemUsloviiaFiltra(ctx context.Context, args CatalogFiltryDliaElektronnykhPisemUsloviiaFiltraArgs) (*CatalogFiltryDliaElektronnykhPisemUsloviiaFiltra, error) {
	return r.Client.CatalogFiltryDliaElektronnykhPisemUsloviiaFiltra(args.Key, nil)
}
func (r *GqlResolver) CatalogFiltryDliaElektronnykhPisemUsloviiaFiltras(ctx context.Context, args CatalogFiltryDliaElektronnykhPisemUsloviiaFiltrasArgs) (*[]CatalogFiltryDliaElektronnykhPisemUsloviiaFiltra, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogFiltryDliaElektronnykhPisemUsloviiaFiltras(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogFiltryDliaElektronnykhPisemUsloviiaFiltras(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogFiltryDliaElektronnykhPisemUsloviiaFiltras(Where{})
}
func (r *GqlResolver) DocumentPreiskurantTsenNaTsvKamni(ctx context.Context, args DocumentPreiskurantTsenNaTsvKamniArgs) (*DocumentPreiskurantTsenNaTsvKamni, error) {
	return r.Client.DocumentPreiskurantTsenNaTsvKamni(args.Key, nil)
}
func (r *GqlResolver) DocumentPreiskurantTsenNaTsvKamnis(ctx context.Context, args DocumentPreiskurantTsenNaTsvKamnisArgs) (*[]DocumentPreiskurantTsenNaTsvKamni, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPreiskurantTsenNaTsvKamnis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPreiskurantTsenNaTsvKamnis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPreiskurantTsenNaTsvKamnis(Where{})
}
func (r *GqlResolver) DocumentPreiskurantTsenNaTsvKamniTablitsa(ctx context.Context, args DocumentPreiskurantTsenNaTsvKamniTablitsaArgs) (*DocumentPreiskurantTsenNaTsvKamniTablitsa, error) {
	return r.Client.DocumentPreiskurantTsenNaTsvKamniTablitsa(args.Key, nil)
}
func (r *GqlResolver) DocumentPreiskurantTsenNaTsvKamniTablitsas(ctx context.Context, args DocumentPreiskurantTsenNaTsvKamniTablitsasArgs) (*[]DocumentPreiskurantTsenNaTsvKamniTablitsa, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPreiskurantTsenNaTsvKamniTablitsas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPreiskurantTsenNaTsvKamniTablitsas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPreiskurantTsenNaTsvKamniTablitsas(Where{})
}
func (r *GqlResolver) Size(ctx context.Context, args SizeArgs) (*Size, error) {
	return r.Client.Size(args.Key, nil)
}
func (r *GqlResolver) Sizes(ctx context.Context, args SizesArgs) (*[]Size, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Sizes(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Sizes(Where{Filter: *args.Filter})
	}
	return r.Client.Sizes(Where{})
}
func (r *GqlResolver) CatalogTipyDragotsennykhMetallov(ctx context.Context, args CatalogTipyDragotsennykhMetallovArgs) (*CatalogTipyDragotsennykhMetallov, error) {
	return r.Client.CatalogTipyDragotsennykhMetallov(args.Key, nil)
}
func (r *GqlResolver) CatalogTipyDragotsennykhMetallovs(ctx context.Context, args CatalogTipyDragotsennykhMetallovsArgs) (*[]CatalogTipyDragotsennykhMetallov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTipyDragotsennykhMetallovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTipyDragotsennykhMetallovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTipyDragotsennykhMetallovs(Where{})
}
func (r *GqlResolver) DocumentTelemarketing(ctx context.Context, args DocumentTelemarketingArgs) (*DocumentTelemarketing, error) {
	return r.Client.DocumentTelemarketing(args.Key, nil)
}
func (r *GqlResolver) DocumentTelemarketings(ctx context.Context, args DocumentTelemarketingsArgs) (*[]DocumentTelemarketing, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentTelemarketings(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentTelemarketings(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentTelemarketings(Where{})
}
func (r *GqlResolver) DocumentTelemarketingUchastniki(ctx context.Context, args DocumentTelemarketingUchastnikiArgs) (*DocumentTelemarketingUchastniki, error) {
	return r.Client.DocumentTelemarketingUchastniki(args.Key, nil)
}
func (r *GqlResolver) DocumentTelemarketingUchastnikis(ctx context.Context, args DocumentTelemarketingUchastnikisArgs) (*[]DocumentTelemarketingUchastniki, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentTelemarketingUchastnikis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentTelemarketingUchastnikis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentTelemarketingUchastnikis(Where{})
}
func (r *GqlResolver) DocumentVozvratDavalcheskogoMetalla(ctx context.Context, args DocumentVozvratDavalcheskogoMetallaArgs) (*DocumentVozvratDavalcheskogoMetalla, error) {
	return r.Client.DocumentVozvratDavalcheskogoMetalla(args.Key, nil)
}
func (r *GqlResolver) DocumentVozvratDavalcheskogoMetallas(ctx context.Context, args DocumentVozvratDavalcheskogoMetallasArgs) (*[]DocumentVozvratDavalcheskogoMetalla, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentVozvratDavalcheskogoMetallas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentVozvratDavalcheskogoMetallas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentVozvratDavalcheskogoMetallas(Where{})
}
func (r *GqlResolver) CatalogAdresnyeSokrashcheniia(ctx context.Context, args CatalogAdresnyeSokrashcheniiaArgs) (*CatalogAdresnyeSokrashcheniia, error) {
	return r.Client.CatalogAdresnyeSokrashcheniia(args.Key, nil)
}
func (r *GqlResolver) CatalogAdresnyeSokrashcheniias(ctx context.Context, args CatalogAdresnyeSokrashcheniiasArgs) (*[]CatalogAdresnyeSokrashcheniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogAdresnyeSokrashcheniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogAdresnyeSokrashcheniias(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogAdresnyeSokrashcheniias(Where{})
}
func (r *GqlResolver) DocumentRassylkaAnket(ctx context.Context, args DocumentRassylkaAnketArgs) (*DocumentRassylkaAnket, error) {
	return r.Client.DocumentRassylkaAnket(args.Key, nil)
}
func (r *GqlResolver) DocumentRassylkaAnkets(ctx context.Context, args DocumentRassylkaAnketsArgs) (*[]DocumentRassylkaAnket, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRassylkaAnkets(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRassylkaAnkets(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRassylkaAnkets(Where{})
}
func (r *GqlResolver) DocumentRassylkaAnketVlozheniia(ctx context.Context, args DocumentRassylkaAnketVlozheniiaArgs) (*DocumentRassylkaAnketVlozheniia, error) {
	return r.Client.DocumentRassylkaAnketVlozheniia(args.Key, nil)
}
func (r *GqlResolver) DocumentRassylkaAnketVlozheniias(ctx context.Context, args DocumentRassylkaAnketVlozheniiasArgs) (*[]DocumentRassylkaAnketVlozheniia, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRassylkaAnketVlozheniias(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRassylkaAnketVlozheniias(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRassylkaAnketVlozheniias(Where{})
}
func (r *GqlResolver) DocumentRassylkaAnketPoluchateli(ctx context.Context, args DocumentRassylkaAnketPoluchateliArgs) (*DocumentRassylkaAnketPoluchateli, error) {
	return r.Client.DocumentRassylkaAnketPoluchateli(args.Key, nil)
}
func (r *GqlResolver) DocumentRassylkaAnketPoluchatelis(ctx context.Context, args DocumentRassylkaAnketPoluchatelisArgs) (*[]DocumentRassylkaAnketPoluchateli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentRassylkaAnketPoluchatelis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentRassylkaAnketPoluchatelis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentRassylkaAnketPoluchatelis(Where{})
}
func (r *GqlResolver) CatalogVidyDeiatelnostiKontragentov(ctx context.Context, args CatalogVidyDeiatelnostiKontragentovArgs) (*CatalogVidyDeiatelnostiKontragentov, error) {
	return r.Client.CatalogVidyDeiatelnostiKontragentov(args.Key, nil)
}
func (r *GqlResolver) CatalogVidyDeiatelnostiKontragentovs(ctx context.Context, args CatalogVidyDeiatelnostiKontragentovsArgs) (*[]CatalogVidyDeiatelnostiKontragentov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogVidyDeiatelnostiKontragentovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogVidyDeiatelnostiKontragentovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogVidyDeiatelnostiKontragentovs(Where{})
}
func (r *GqlResolver) CatalogTorgovoeOborudovanie(ctx context.Context, args CatalogTorgovoeOborudovanieArgs) (*CatalogTorgovoeOborudovanie, error) {
	return r.Client.CatalogTorgovoeOborudovanie(args.Key, nil)
}
func (r *GqlResolver) CatalogTorgovoeOborudovanies(ctx context.Context, args CatalogTorgovoeOborudovaniesArgs) (*[]CatalogTorgovoeOborudovanie, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTorgovoeOborudovanies(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTorgovoeOborudovanies(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTorgovoeOborudovanies(Where{})
}
func (r *GqlResolver) CatalogSkhemyRealizatsii(ctx context.Context, args CatalogSkhemyRealizatsiiArgs) (*CatalogSkhemyRealizatsii, error) {
	return r.Client.CatalogSkhemyRealizatsii(args.Key, nil)
}
func (r *GqlResolver) CatalogSkhemyRealizatsiis(ctx context.Context, args CatalogSkhemyRealizatsiisArgs) (*[]CatalogSkhemyRealizatsii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSkhemyRealizatsiis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSkhemyRealizatsiis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSkhemyRealizatsiis(Where{})
}
func (r *GqlResolver) CatalogSkhemyRealizatsiiEtapySkhemy(ctx context.Context, args CatalogSkhemyRealizatsiiEtapySkhemyArgs) (*CatalogSkhemyRealizatsiiEtapySkhemy, error) {
	return r.Client.CatalogSkhemyRealizatsiiEtapySkhemy(args.Key, nil)
}
func (r *GqlResolver) CatalogSkhemyRealizatsiiEtapySkhemys(ctx context.Context, args CatalogSkhemyRealizatsiiEtapySkhemysArgs) (*[]CatalogSkhemyRealizatsiiEtapySkhemy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogSkhemyRealizatsiiEtapySkhemys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogSkhemyRealizatsiiEtapySkhemys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogSkhemyRealizatsiiEtapySkhemys(Where{})
}
func (r *GqlResolver) CatalogPodkliuchaemoeOborudovanie(ctx context.Context, args CatalogPodkliuchaemoeOborudovanieArgs) (*CatalogPodkliuchaemoeOborudovanie, error) {
	return r.Client.CatalogPodkliuchaemoeOborudovanie(args.Key, nil)
}
func (r *GqlResolver) CatalogPodkliuchaemoeOborudovanies(ctx context.Context, args CatalogPodkliuchaemoeOborudovaniesArgs) (*[]CatalogPodkliuchaemoeOborudovanie, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogPodkliuchaemoeOborudovanies(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogPodkliuchaemoeOborudovanies(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogPodkliuchaemoeOborudovanies(Where{})
}
func (r *GqlResolver) DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii(ctx context.Context, args DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiArgs) (*DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii, error) {
	return r.Client.DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii(args.Key, nil)
}
func (r *GqlResolver) DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniis(ctx context.Context, args DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniisArgs) (*[]DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnoshenii, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniis(Where{})
}
func (r *GqlResolver) DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov(ctx context.Context, args DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovArgs) (*DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov, error) {
	return r.Client.DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov(args.Key, nil)
}
func (r *GqlResolver) DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovs(ctx context.Context, args DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovsArgs) (*[]DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentKlassifikatsiiaPokupateleiPoStadiiamVzaimootnosheniiTablitsaRaspredeleniiaKontragentovs(Where{})
}
func (r *GqlResolver) CatalogGabarity(ctx context.Context, args CatalogGabarityArgs) (*CatalogGabarity, error) {
	return r.Client.CatalogGabarity(args.Key, nil)
}
func (r *GqlResolver) CatalogGabaritys(ctx context.Context, args CatalogGabaritysArgs) (*[]CatalogGabarity, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGabaritys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGabaritys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGabaritys(Where{})
}
func (r *GqlResolver) DocumentZakazKlienta(ctx context.Context, args DocumentZakazKlientaArgs) (*DocumentZakazKlienta, error) {
	return r.Client.DocumentZakazKlienta(args.Key, nil)
}
func (r *GqlResolver) DocumentZakazKlientas(ctx context.Context, args DocumentZakazKlientasArgs) (*[]DocumentZakazKlienta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakazKlientas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakazKlientas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakazKlientas(Where{})
}
func (r *GqlResolver) DocumentZakazKlientaTovary(ctx context.Context, args DocumentZakazKlientaTovaryArgs) (*DocumentZakazKlientaTovary, error) {
	return r.Client.DocumentZakazKlientaTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentZakazKlientaTovarys(ctx context.Context, args DocumentZakazKlientaTovarysArgs) (*[]DocumentZakazKlientaTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentZakazKlientaTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentZakazKlientaTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentZakazKlientaTovarys(Where{})
}
func (r *GqlResolver) ArriveFromManufacturing(ctx context.Context, args ArriveFromManufacturingArgs) (*ArriveFromManufacturing, error) {
	return r.Client.ArriveFromManufacturing(args.Key, nil)
}
func (r *GqlResolver) ArriveFromManufacturings(ctx context.Context, args ArriveFromManufacturingsArgs) (*[]ArriveFromManufacturing, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.ArriveFromManufacturings(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.ArriveFromManufacturings(Where{Filter: *args.Filter})
	}
	return r.Client.ArriveFromManufacturings(Where{})
}
func (r *GqlResolver) ArriveFromManufacturingInstance(ctx context.Context, args ArriveFromManufacturingInstanceArgs) (*ArriveFromManufacturingInstance, error) {
	return r.Client.ArriveFromManufacturingInstance(args.Key, nil)
}
func (r *GqlResolver) ArriveFromManufacturingInstances(ctx context.Context, args ArriveFromManufacturingInstancesArgs) (*[]ArriveFromManufacturingInstance, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.ArriveFromManufacturingInstances(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.ArriveFromManufacturingInstances(Where{Filter: *args.Filter})
	}
	return r.Client.ArriveFromManufacturingInstances(Where{})
}
func (r *GqlResolver) DocumentPostuplenieProduktsiiIzProizvodstvaMaterialy(ctx context.Context, args DocumentPostuplenieProduktsiiIzProizvodstvaMaterialyArgs) (*DocumentPostuplenieProduktsiiIzProizvodstvaMaterialy, error) {
	return r.Client.DocumentPostuplenieProduktsiiIzProizvodstvaMaterialy(args.Key, nil)
}
func (r *GqlResolver) DocumentPostuplenieProduktsiiIzProizvodstvaMaterialys(ctx context.Context, args DocumentPostuplenieProduktsiiIzProizvodstvaMaterialysArgs) (*[]DocumentPostuplenieProduktsiiIzProizvodstvaMaterialy, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPostuplenieProduktsiiIzProizvodstvaMaterialys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPostuplenieProduktsiiIzProizvodstvaMaterialys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPostuplenieProduktsiiIzProizvodstvaMaterialys(Where{})
}
func (r *GqlResolver) DocumentJournalZakazyPostavshchikam(ctx context.Context, args DocumentJournalZakazyPostavshchikamArgs) (*DocumentJournalZakazyPostavshchikam, error) {
	return r.Client.DocumentJournalZakazyPostavshchikam(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalZakazyPostavshchikams(ctx context.Context, args DocumentJournalZakazyPostavshchikamsArgs) (*[]DocumentJournalZakazyPostavshchikam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalZakazyPostavshchikams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalZakazyPostavshchikams(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalZakazyPostavshchikams(Where{})
}
func (r *GqlResolver) DocumentJournalSkladskieDokumenty(ctx context.Context, args DocumentJournalSkladskieDokumentyArgs) (*DocumentJournalSkladskieDokumenty, error) {
	return r.Client.DocumentJournalSkladskieDokumenty(args.Key, nil)
}
func (r *GqlResolver) DocumentJournalSkladskieDokumentys(ctx context.Context, args DocumentJournalSkladskieDokumentysArgs) (*[]DocumentJournalSkladskieDokumenty, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentJournalSkladskieDokumentys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentJournalSkladskieDokumentys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentJournalSkladskieDokumentys(Where{})
}
func (r *GqlResolver) CatalogsmsUsloviiaOtboraDiskontnykhKart(ctx context.Context, args CatalogsmsUsloviiaOtboraDiskontnykhKartArgs) (*CatalogsmsUsloviiaOtboraDiskontnykhKart, error) {
	return r.Client.CatalogsmsUsloviiaOtboraDiskontnykhKart(args.Key, nil)
}
func (r *GqlResolver) CatalogsmsUsloviiaOtboraDiskontnykhKarts(ctx context.Context, args CatalogsmsUsloviiaOtboraDiskontnykhKartsArgs) (*[]CatalogsmsUsloviiaOtboraDiskontnykhKart, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogsmsUsloviiaOtboraDiskontnykhKarts(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogsmsUsloviiaOtboraDiskontnykhKarts(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogsmsUsloviiaOtboraDiskontnykhKarts(Where{})
}
func (r *GqlResolver) Arrive(ctx context.Context, args ArriveArgs) (*Arrive, error) {
	return r.Client.Arrive(args.Key, nil)
}
func (r *GqlResolver) Arrives(ctx context.Context, args ArrivesArgs) (*[]Arrive, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.Arrives(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.Arrives(Where{Filter: *args.Filter})
	}
	return r.Client.Arrives(Where{})
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugTovary(ctx context.Context, args DocumentPostuplenieTovarovUslugTovaryArgs) (*DocumentPostuplenieTovarovUslugTovary, error) {
	return r.Client.DocumentPostuplenieTovarovUslugTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugTovarys(ctx context.Context, args DocumentPostuplenieTovarovUslugTovarysArgs) (*[]DocumentPostuplenieTovarovUslugTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPostuplenieTovarovUslugTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPostuplenieTovarovUslugTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPostuplenieTovarovUslugTovarys(Where{})
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugUslugi(ctx context.Context, args DocumentPostuplenieTovarovUslugUslugiArgs) (*DocumentPostuplenieTovarovUslugUslugi, error) {
	return r.Client.DocumentPostuplenieTovarovUslugUslugi(args.Key, nil)
}
func (r *GqlResolver) DocumentPostuplenieTovarovUslugUslugis(ctx context.Context, args DocumentPostuplenieTovarovUslugUslugisArgs) (*[]DocumentPostuplenieTovarovUslugUslugi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPostuplenieTovarovUslugUslugis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPostuplenieTovarovUslugUslugis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPostuplenieTovarovUslugUslugis(Where{})
}
func (r *GqlResolver) DocumentSchetFakturaVydannyi(ctx context.Context, args DocumentSchetFakturaVydannyiArgs) (*DocumentSchetFakturaVydannyi, error) {
	return r.Client.DocumentSchetFakturaVydannyi(args.Key, nil)
}
func (r *GqlResolver) DocumentSchetFakturaVydannyis(ctx context.Context, args DocumentSchetFakturaVydannyisArgs) (*[]DocumentSchetFakturaVydannyi, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentSchetFakturaVydannyis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentSchetFakturaVydannyis(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentSchetFakturaVydannyis(Where{})
}
func (r *GqlResolver) DocumentPlanProdazhPoSalonam(ctx context.Context, args DocumentPlanProdazhPoSalonamArgs) (*DocumentPlanProdazhPoSalonam, error) {
	return r.Client.DocumentPlanProdazhPoSalonam(args.Key, nil)
}
func (r *GqlResolver) DocumentPlanProdazhPoSalonams(ctx context.Context, args DocumentPlanProdazhPoSalonamsArgs) (*[]DocumentPlanProdazhPoSalonam, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlanProdazhPoSalonams(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlanProdazhPoSalonams(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlanProdazhPoSalonams(Where{})
}
func (r *GqlResolver) DocumentPlanProdazhPoSalonamSalony(ctx context.Context, args DocumentPlanProdazhPoSalonamSalonyArgs) (*DocumentPlanProdazhPoSalonamSalony, error) {
	return r.Client.DocumentPlanProdazhPoSalonamSalony(args.Key, nil)
}
func (r *GqlResolver) DocumentPlanProdazhPoSalonamSalonys(ctx context.Context, args DocumentPlanProdazhPoSalonamSalonysArgs) (*[]DocumentPlanProdazhPoSalonamSalony, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlanProdazhPoSalonamSalonys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlanProdazhPoSalonamSalonys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlanProdazhPoSalonamSalonys(Where{})
}
func (r *GqlResolver) DocumentPlanProdazhPoSalonamDniPoGrafiku(ctx context.Context, args DocumentPlanProdazhPoSalonamDniPoGrafikuArgs) (*DocumentPlanProdazhPoSalonamDniPoGrafiku, error) {
	return r.Client.DocumentPlanProdazhPoSalonamDniPoGrafiku(args.Key, nil)
}
func (r *GqlResolver) DocumentPlanProdazhPoSalonamDniPoGrafikus(ctx context.Context, args DocumentPlanProdazhPoSalonamDniPoGrafikusArgs) (*[]DocumentPlanProdazhPoSalonamDniPoGrafiku, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPlanProdazhPoSalonamDniPoGrafikus(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPlanProdazhPoSalonamDniPoGrafikus(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPlanProdazhPoSalonamDniPoGrafikus(Where{})
}
func (r *GqlResolver) CatalogBankovskieScheta(ctx context.Context, args CatalogBankovskieSchetaArgs) (*CatalogBankovskieScheta, error) {
	return r.Client.CatalogBankovskieScheta(args.Key, nil)
}
func (r *GqlResolver) CatalogBankovskieSchetas(ctx context.Context, args CatalogBankovskieSchetasArgs) (*[]CatalogBankovskieScheta, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogBankovskieSchetas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogBankovskieSchetas(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogBankovskieSchetas(Where{})
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomitentuOProdazhakh(ctx context.Context, args DocumentStornirovanieOtchetaKomitentuOProdazhakhArgs) (*DocumentStornirovanieOtchetaKomitentuOProdazhakh, error) {
	return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakh(args.Key, nil)
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomitentuOProdazhakhs(ctx context.Context, args DocumentStornirovanieOtchetaKomitentuOProdazhakhsArgs) (*[]DocumentStornirovanieOtchetaKomitentuOProdazhakh, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhs(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhs(Where{})
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva(ctx context.Context, args DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvaArgs) (*DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva, error) {
	return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva(args.Key, nil)
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvas(ctx context.Context, args DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvasArgs) (*[]DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstva, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvas(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvas(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhDenezhnyeSredstvas(Where{})
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomitentuOProdazhakhTovary(ctx context.Context, args DocumentStornirovanieOtchetaKomitentuOProdazhakhTovaryArgs) (*DocumentStornirovanieOtchetaKomitentuOProdazhakhTovary, error) {
	return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentStornirovanieOtchetaKomitentuOProdazhakhTovarys(ctx context.Context, args DocumentStornirovanieOtchetaKomitentuOProdazhakhTovarysArgs) (*[]DocumentStornirovanieOtchetaKomitentuOProdazhakhTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentStornirovanieOtchetaKomitentuOProdazhakhTovarys(Where{})
}
func (r *GqlResolver) DocumentPeredachaVRemont(ctx context.Context, args DocumentPeredachaVRemontArgs) (*DocumentPeredachaVRemont, error) {
	return r.Client.DocumentPeredachaVRemont(args.Key, nil)
}
func (r *GqlResolver) DocumentPeredachaVRemonts(ctx context.Context, args DocumentPeredachaVRemontsArgs) (*[]DocumentPeredachaVRemont, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPeredachaVRemonts(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPeredachaVRemonts(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPeredachaVRemonts(Where{})
}
func (r *GqlResolver) DocumentPeredachaVRemontIzdeliiaPriniatyeVRemont(ctx context.Context, args DocumentPeredachaVRemontIzdeliiaPriniatyeVRemontArgs) (*DocumentPeredachaVRemontIzdeliiaPriniatyeVRemont, error) {
	return r.Client.DocumentPeredachaVRemontIzdeliiaPriniatyeVRemont(args.Key, nil)
}
func (r *GqlResolver) DocumentPeredachaVRemontIzdeliiaPriniatyeVRemonts(ctx context.Context, args DocumentPeredachaVRemontIzdeliiaPriniatyeVRemontsArgs) (*[]DocumentPeredachaVRemontIzdeliiaPriniatyeVRemont, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPeredachaVRemontIzdeliiaPriniatyeVRemonts(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPeredachaVRemontIzdeliiaPriniatyeVRemonts(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPeredachaVRemontIzdeliiaPriniatyeVRemonts(Where{})
}
func (r *GqlResolver) DocumentPeredachaVRemontTovary(ctx context.Context, args DocumentPeredachaVRemontTovaryArgs) (*DocumentPeredachaVRemontTovary, error) {
	return r.Client.DocumentPeredachaVRemontTovary(args.Key, nil)
}
func (r *GqlResolver) DocumentPeredachaVRemontTovarys(ctx context.Context, args DocumentPeredachaVRemontTovarysArgs) (*[]DocumentPeredachaVRemontTovary, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.DocumentPeredachaVRemontTovarys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.DocumentPeredachaVRemontTovarys(Where{Filter: *args.Filter})
	}
	return r.Client.DocumentPeredachaVRemontTovarys(Where{})
}
func (r *GqlResolver) CatalogPolzovateli(ctx context.Context, args CatalogPolzovateliArgs) (*CatalogPolzovateli, error) {
	return r.Client.CatalogPolzovateli(args.Key, nil)
}
func (r *GqlResolver) CatalogPolzovatelis(ctx context.Context, args CatalogPolzovatelisArgs) (*[]CatalogPolzovateli, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogPolzovatelis(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogPolzovatelis(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogPolzovatelis(Where{})
}
func (r *GqlResolver) CatalogTsenovyeKoridory(ctx context.Context, args CatalogTsenovyeKoridoryArgs) (*CatalogTsenovyeKoridory, error) {
	return r.Client.CatalogTsenovyeKoridory(args.Key, nil)
}
func (r *GqlResolver) CatalogTsenovyeKoridorys(ctx context.Context, args CatalogTsenovyeKoridorysArgs) (*[]CatalogTsenovyeKoridory, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogTsenovyeKoridorys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogTsenovyeKoridorys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogTsenovyeKoridorys(Where{})
}
func (r *GqlResolver) CatalogGruppySkladov(ctx context.Context, args CatalogGruppySkladovArgs) (*CatalogGruppySkladov, error) {
	return r.Client.CatalogGruppySkladov(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppySkladovs(ctx context.Context, args CatalogGruppySkladovsArgs) (*[]CatalogGruppySkladov, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppySkladovs(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppySkladovs(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppySkladovs(Where{})
}
func (r *GqlResolver) CatalogGruppySkladovSklady(ctx context.Context, args CatalogGruppySkladovSkladyArgs) (*CatalogGruppySkladovSklady, error) {
	return r.Client.CatalogGruppySkladovSklady(args.Key, nil)
}
func (r *GqlResolver) CatalogGruppySkladovSkladys(ctx context.Context, args CatalogGruppySkladovSkladysArgs) (*[]CatalogGruppySkladovSklady, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.CatalogGruppySkladovSkladys(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.CatalogGruppySkladovSkladys(Where{Filter: *args.Filter})
	}
	return r.Client.CatalogGruppySkladovSkladys(Where{})
}
