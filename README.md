**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-13 13:59:33 UTC（2026-03-13 21:59:33 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1735ms | 否 | ✓ 927ms | 否 | ✓ 1049ms | http |
| 114.214.208.153:10808 | 否 | ✓ 1828ms | ✓ 1657ms | ✓ 1649ms | ✓ 1233ms | http |
| 45.167.124.52:8080 | ✓ 1832ms | 否 | ✓ 1373ms | 否 | ✓ 1710ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1800ms | ✓ 1370ms | ✓ 1230ms | ✓ 962ms | http |
| 20.27.13.35:8561 | 否 | ✓ 1256ms | ✓ 474ms | ✓ 982ms | ✓ 1660ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1749ms | ✓ 1510ms | ✓ 1168ms | http |
| 101.43.255.96:80 | 否 | 否 | ✓ 1276ms | ✓ 1515ms | ✓ 1654ms | http |
| 14.225.211.139:7890 | 否 | ✓ 1758ms | 否 | ✓ 1072ms | ✓ 1037ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1215ms | 否 | ✓ 1458ms | ✓ 951ms | http |
| 81.70.169.194:80 | ✓ 1676ms | ✓ 1302ms | 否 | ✓ 1219ms | ✓ 1874ms | http |
| 45.136.130.223:8443 | ✓ 208ms | ✓ 663ms | ✓ 269ms | ✓ 696ms | ✓ 515ms | http |
| 117.159.239.49:22222 | ✓ 1032ms | ✓ 1506ms | ✓ 897ms | ✓ 1567ms | ✓ 867ms | http |
| 113.59.32.160:22222 | ✓ 1135ms | ✓ 1401ms | ✓ 1427ms | ✓ 1354ms | ✓ 1053ms | http |
| 120.238.159.189:22222 | ✓ 1466ms | ✓ 1758ms | ✓ 1175ms | ✓ 1817ms | 否 | http |
| 120.238.159.230:22222 | ✓ 1825ms | ✓ 1954ms | ✓ 1104ms | ✓ 1872ms | 否 | http |
| 183.249.5.109:22222 | ✓ 778ms | ✓ 911ms | ✓ 778ms | ✓ 983ms | ✓ 777ms | http |
| 62.60.177.204:34094 | 否 | 否 | ✓ 1346ms | ✓ 1123ms | ✓ 1041ms | http |
| 117.159.239.44:22222 | 否 | ✓ 1025ms | ✓ 1143ms | ✓ 1072ms | ✓ 875ms | http |
| 120.232.242.119:22222 | ✓ 1024ms | 否 | 否 | ✓ 1146ms | ✓ 1042ms | http |
| 117.159.239.51:22222 | 否 | ✓ 1080ms | ✓ 1886ms | ✓ 1085ms | ✓ 823ms | http |
| 117.159.239.52:22222 | 否 | 否 | ✓ 913ms | ✓ 1106ms | ✓ 863ms | http |
| 152.42.213.210:443 | ✓ 1026ms | 否 | ✓ 1056ms | 否 | ✓ 891ms | http |
| 152.42.213.210:8080 | ✓ 738ms | 否 | ✓ 1911ms | 否 | ✓ 991ms | http |
| 120.238.159.228:22222 | ✓ 1744ms | 否 | ✓ 1818ms | ✓ 1458ms | ✓ 1063ms | http |
| 222.184.48.251:22222 | ✓ 1015ms | 否 | ✓ 1659ms | ✓ 1474ms | 否 | http |
| 45.136.198.40:3128 | 否 | 否 | ✓ 1728ms | ✓ 1641ms | ✓ 1274ms | http |
| 152.53.194.38:7890 | 否 | 否 | ✓ 1477ms | ✓ 1786ms | ✓ 1541ms | http |
| 113.59.32.148:22222 | ✓ 1468ms | ✓ 1792ms | ✓ 1265ms | ✓ 1158ms | ✓ 981ms | http |
| 113.59.32.142:22222 | 否 | ✓ 1890ms | ✓ 1069ms | ✓ 1449ms | ✓ 1231ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1835ms | 否 | ✓ 1459ms | ✓ 1970ms | http |
| 178.236.245.17:3128 | 否 | 否 | ✓ 1258ms | ✓ 1805ms | ✓ 1841ms | http |
| 20.27.15.49:8561 | 否 | 否 | ✓ 569ms | ✓ 950ms | ✓ 852ms | http |
| 89.251.9.11:3128 | 否 | 否 | ✓ 1801ms | ✓ 1447ms | ✓ 1100ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1449ms | 否 | ✓ 1142ms | ✓ 1045ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1216ms | ✓ 1916ms | ✓ 1819ms | http |
| 8.222.175.80:6128 | 否 | 否 | ✓ 893ms | ✓ 1082ms | ✓ 848ms | http |
| 137.184.14.135:3128 | 否 | 否 | ✓ 822ms | ✓ 745ms | ✓ 548ms | http |
| 207.254.71.62:8088 | ✓ 1125ms | 否 | ✓ 1346ms | ✓ 1688ms | ✓ 1518ms | http |
| 106.117.208.101:7890 | ✓ 1210ms | 否 | ✓ 1163ms | 否 | ✓ 1537ms | http |
| 45.186.6.104:3128 | ✓ 1348ms | ✓ 1645ms | ✓ 1725ms | 否 | 否 | http |
| 120.240.35.173:22222 | ✓ 996ms | 否 | ✓ 1041ms | ✓ 1159ms | ✓ 937ms | http |
| 46.183.25.8:443 | ✓ 1823ms | 否 | ✓ 1064ms | ✓ 1712ms | 否 | http |
| 82.165.61.217:8085 | ✓ 1191ms | 否 | ✓ 1407ms | 否 | ✓ 1897ms | http |
| 14.225.212.37:7890 | ✓ 896ms | ✓ 1937ms | ✓ 861ms | ✓ 1067ms | ✓ 848ms | http |
| 47.77.193.180:1080 | 否 | 否 | ✓ 1373ms | ✓ 984ms | ✓ 1727ms | http |
| 211.171.114.154:3128 | 否 | 否 | ✓ 1697ms | ✓ 1518ms | ✓ 1438ms | http |
| 45.136.131.47:8443 | ✓ 1730ms | ✓ 781ms | ✓ 942ms | ✓ 782ms | ✓ 669ms | http |
| 45.136.130.191:8443 | ✓ 450ms | ✓ 989ms | ✓ 95ms | ✓ 784ms | ✓ 516ms | http |
| 45.136.130.188:8443 | ✓ 1129ms | ✓ 624ms | ✓ 88ms | ✓ 688ms | ✓ 496ms | http |
| 183.249.5.110:22222 | ✓ 901ms | ✓ 925ms | ✓ 679ms | ✓ 1078ms | 否 | http |
| 183.249.5.111:22222 | ✓ 1924ms | ✓ 873ms | ✓ 678ms | ✓ 894ms | 否 | http |
| 113.59.32.141:22222 | ✓ 1065ms | ✓ 1409ms | ✓ 1247ms | ✓ 1457ms | ✓ 1057ms | http |
| 120.240.35.176:22222 | ✓ 1939ms | ✓ 1243ms | ✓ 1175ms | ✓ 1099ms | ✓ 913ms | http |
| 120.240.35.177:22222 | 否 | 否 | ✓ 1036ms | ✓ 1209ms | ✓ 931ms | http |
| 120.240.35.160:22222 | 否 | ✓ 1204ms | ✓ 1317ms | 否 | ✓ 1000ms | http |
| 20.27.14.220:8561 | ✓ 1459ms | ✓ 756ms | ✓ 573ms | ✓ 776ms | ✓ 596ms | http |
| 183.249.5.214:22222 | 否 | ✓ 896ms | ✓ 1024ms | 否 | ✓ 1034ms | http |
| 20.78.26.206:8561 | ✓ 1479ms | ✓ 1328ms | ✓ 457ms | ✓ 812ms | ✓ 693ms | http |
| 20.210.39.153:8561 | ✓ 1942ms | ✓ 1889ms | ✓ 455ms | ✓ 863ms | ✓ 642ms | http |
| 20.27.11.248:8561 | ✓ 1958ms | ✓ 1879ms | ✓ 445ms | ✓ 865ms | ✓ 643ms | http |
| 222.184.48.235:22222 | 否 | 否 | ✓ 1729ms | ✓ 1316ms | ✓ 1562ms | http |
| 20.78.118.91:8561 | ✓ 536ms | ✓ 1908ms | ✓ 453ms | ✓ 1930ms | ✓ 632ms | http |
| 117.159.239.46:22222 | ✓ 1480ms | 否 | ✓ 1708ms | ✓ 1921ms | ✓ 835ms | http |
| 120.240.29.173:22222 | 否 | 否 | ✓ 939ms | ✓ 1370ms | ✓ 919ms | http |
| 103.113.70.189:1081 | ✓ 771ms | 否 | 否 | ✓ 1286ms | ✓ 962ms | http |
| 220.197.44.36:3128 | ✓ 1581ms | 否 | ✓ 1299ms | 否 | ✓ 1932ms | http |
| 103.84.95.54:7890 | ✓ 894ms | 否 | ✓ 917ms | ✓ 959ms | 否 | http |
| 101.47.73.135:3128 | ✓ 826ms | 否 | ✓ 871ms | ✓ 1000ms | ✓ 1025ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
