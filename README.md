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

最后更新：2026-03-30 08:46:48 UTC（2026-03-30 16:46:48 UTC+8）

**代理总数：117**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 117 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 117 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 39.185.46.193:5911 | ✓ 815ms | ✓ 799ms | ✓ 696ms | ✓ 1189ms | ✓ 742ms | http |
| 35.225.22.61:80 | ✓ 447ms | 否 | ✓ 1063ms | ✓ 1279ms | ✓ 1181ms | http |
| 43.99.54.236:5555 | ✓ 1286ms | ✓ 1763ms | ✓ 626ms | ✓ 1231ms | ✓ 642ms | http |
| 167.103.115.102:8800 | ✓ 1061ms | 否 | ✓ 922ms | ✓ 990ms | ✓ 1195ms | http |
| 147.161.210.140:8800 | ✓ 1416ms | 否 | ✓ 844ms | ✓ 1372ms | ✓ 1276ms | http |
| 1.231.81.166:3128 | ✓ 1446ms | ✓ 1158ms | ✓ 1436ms | 否 | ✓ 816ms | http |
| 113.160.132.26:8080 | ✓ 1429ms | ✓ 1628ms | ✓ 875ms | ✓ 1296ms | ✓ 1043ms | http |
| 147.161.239.240:8800 | ✓ 1199ms | ✓ 1951ms | ✓ 988ms | ✓ 1646ms | ✓ 1430ms | http |
| 42.96.16.158:1311 | ✓ 1427ms | 否 | ✓ 1232ms | ✓ 1226ms | 否 | http |
| 95.213.217.168:52004 | ✓ 1232ms | ✓ 1799ms | ✓ 1283ms | 否 | ✓ 1761ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 740ms | ✓ 842ms | ✓ 952ms | http |
| 167.103.34.108:8800 | ✓ 1271ms | 否 | ✓ 1114ms | ✓ 1298ms | 否 | http |
| 101.47.73.135:3128 | ✓ 802ms | 否 | ✓ 660ms | ✓ 1255ms | 否 | http |
| 59.46.216.131:30001 | 否 | ✓ 1365ms | ✓ 1867ms | ✓ 1545ms | ✓ 1028ms | http |
| 167.103.31.122:8800 | ✓ 1745ms | 否 | ✓ 1804ms | 否 | ✓ 1870ms | http |
| 121.126.185.63:25152 | ✓ 1206ms | ✓ 1557ms | ✓ 1439ms | 否 | 否 | http |
| 167.103.144.127:8800 | ✓ 878ms | 否 | ✓ 810ms | ✓ 1211ms | ✓ 1079ms | http |
| 193.233.22.29:10808 | ✓ 433ms | 否 | ✓ 1229ms | ✓ 1648ms | 否 | http |
| 38.34.179.67:8451 | 否 | ✓ 1053ms | ✓ 684ms | 否 | ✓ 812ms | http |
| 115.231.181.40:8128 | ✓ 1148ms | 否 | ✓ 1937ms | 否 | ✓ 1579ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1054ms | ✓ 1065ms | ✓ 1153ms | ✓ 905ms | http |
| 43.153.28.68:3128 | ✓ 952ms | 否 | ✓ 1427ms | ✓ 677ms | 否 | http |
| 45.88.0.116:3128 | ✓ 664ms | 否 | ✓ 1423ms | ✓ 1595ms | ✓ 1683ms | http |
| 38.34.179.98:8453 | ✓ 161ms | ✓ 610ms | ✓ 313ms | ✓ 844ms | ✓ 522ms | http |
| 38.34.179.34:8445 | ✓ 186ms | ✓ 1185ms | ✓ 338ms | ✓ 783ms | ✓ 518ms | http |
| 38.34.179.39:8445 | ✓ 243ms | ✓ 1783ms | ✓ 281ms | ✓ 670ms | ✓ 548ms | http |
| 208.87.243.199:7878 | ✓ 388ms | ✓ 858ms | ✓ 357ms | ✓ 786ms | ✓ 634ms | http |
| 38.145.208.216:8448 | ✓ 759ms | ✓ 693ms | ✓ 160ms | ✓ 736ms | ✓ 730ms | http |
| 38.34.179.94:8453 | ✓ 797ms | 否 | ✓ 381ms | ✓ 880ms | ✓ 528ms | http |
| 38.145.203.34:8445 | ✓ 981ms | ✓ 659ms | ✓ 107ms | ✓ 807ms | ✓ 601ms | http |
| 38.145.220.27:8446 | ✓ 179ms | ✓ 1341ms | ✓ 491ms | ✓ 1562ms | ✓ 637ms | http |
| 38.34.179.52:8451 | ✓ 153ms | ✓ 1657ms | ✓ 1188ms | ✓ 1743ms | ✓ 493ms | http |
| 38.145.208.218:8448 | ✓ 748ms | 否 | ✓ 620ms | ✓ 1136ms | ✓ 800ms | http |
| 38.34.179.24:8453 | ✓ 144ms | ✓ 1254ms | ✓ 1298ms | ✓ 1793ms | ✓ 576ms | http |
| 38.145.208.220:8447 | ✓ 894ms | ✓ 759ms | ✓ 528ms | ✓ 1871ms | ✓ 1374ms | http |
| 177.234.217.88:999 | ✓ 1401ms | ✓ 1823ms | ✓ 1206ms | ✓ 1877ms | ✓ 1643ms | http |
| 45.88.0.114:3128 | ✓ 1879ms | 否 | ✓ 1775ms | ✓ 1654ms | 否 | http |
| 116.80.65.80:3172 | ✓ 1484ms | 否 | ✓ 1934ms | ✓ 1812ms | 否 | http |
| 121.43.196.210:8222 | ✓ 907ms | ✓ 1050ms | ✓ 852ms | ✓ 1092ms | ✓ 855ms | http |
| 120.92.212.16:7890 | ✓ 917ms | ✓ 1150ms | ✓ 898ms | 否 | ✓ 921ms | http |
| 38.145.220.33:8448 | ✓ 531ms | ✓ 1925ms | ✓ 1037ms | ✓ 840ms | ✓ 516ms | http |
| 38.34.179.161:8448 | ✓ 388ms | ✓ 1326ms | ✓ 1690ms | ✓ 887ms | ✓ 518ms | http |
| 38.34.179.51:8449 | 否 | 否 | ✓ 663ms | ✓ 1197ms | ✓ 1792ms | http |
| 158.160.215.167:8127 | ✓ 1125ms | ✓ 1860ms | ✓ 1556ms | 否 | 否 | http |
| 45.12.151.226:2829 | 否 | 否 | ✓ 1139ms | ✓ 1904ms | ✓ 1405ms | http |
| 89.208.106.138:10808 | ✓ 663ms | ✓ 1844ms | ✓ 1447ms | 否 | 否 | http |
| 38.34.179.100:8449 | ✓ 1533ms | ✓ 1849ms | ✓ 260ms | ✓ 689ms | ✓ 554ms | http |
| 38.145.203.86:8452 | ✓ 1092ms | ✓ 1154ms | ✓ 403ms | ✓ 831ms | ✓ 1134ms | http |
| 45.88.0.111:3128 | ✓ 665ms | 否 | ✓ 900ms | ✓ 1514ms | ✓ 1913ms | http |
| 213.220.62.62:3128 | ✓ 637ms | 否 | ✓ 932ms | ✓ 1520ms | ✓ 1211ms | http |
| 120.92.212.16:8890 | ✓ 886ms | ✓ 1194ms | ✓ 1868ms | ✓ 1212ms | ✓ 1983ms | http |
| 38.34.179.150:8449 | ✓ 837ms | ✓ 753ms | ✓ 95ms | ✓ 688ms | ✓ 500ms | http |
| 45.136.131.50:8452 | ✓ 119ms | ✓ 665ms | ✓ 145ms | ✓ 724ms | ✓ 691ms | http |
| 45.136.131.52:8452 | ✓ 122ms | ✓ 686ms | ✓ 174ms | ✓ 817ms | ✓ 679ms | http |
| 45.136.131.53:8452 | ✓ 147ms | ✓ 776ms | ✓ 145ms | ✓ 830ms | ✓ 716ms | http |
| 45.136.131.54:8448 | ✓ 151ms | ✓ 1725ms | ✓ 102ms | ✓ 757ms | ✓ 701ms | http |
| 38.34.183.222:8453 | ✓ 352ms | ✓ 713ms | ✓ 594ms | ✓ 694ms | ✓ 591ms | http |
| 38.34.183.225:8450 | ✓ 746ms | ✓ 1335ms | ✓ 84ms | ✓ 676ms | ✓ 697ms | http |
| 38.145.220.102:8453 | ✓ 737ms | 否 | ✓ 103ms | ✓ 707ms | ✓ 1048ms | http |
| 8.219.97.248:80 | ✓ 916ms | ✓ 1965ms | ✓ 1265ms | 否 | 否 | http |
| 38.34.179.98:8451 | ✓ 527ms | 否 | ✓ 870ms | ✓ 732ms | ✓ 755ms | http |
| 91.233.223.147:3128 | ✓ 1475ms | 否 | ✓ 1101ms | 否 | ✓ 1843ms | http |
| 45.88.0.115:3128 | 否 | ✓ 1928ms | ✓ 1398ms | ✓ 1857ms | 否 | http |
| 38.145.220.77:8447 | ✓ 586ms | ✓ 916ms | ✓ 550ms | ✓ 1452ms | ✓ 742ms | http |
| 45.88.0.99:3128 | ✓ 791ms | 否 | ✓ 646ms | ✓ 1527ms | ✓ 1185ms | http |
| 45.88.0.98:3128 | ✓ 735ms | ✓ 1639ms | ✓ 705ms | 否 | ✓ 1181ms | http |
| 45.88.0.117:3128 | ✓ 728ms | ✓ 1614ms | ✓ 728ms | 否 | ✓ 1189ms | http |
| 45.88.0.113:3128 | ✓ 667ms | 否 | ✓ 669ms | ✓ 1578ms | ✓ 1188ms | http |
| 217.76.245.80:999 | ✓ 1394ms | ✓ 1369ms | ✓ 1236ms | ✓ 1712ms | ✓ 1551ms | http |
| 101.32.244.83:8080 | ✓ 1576ms | 否 | ✓ 902ms | ✓ 1348ms | ✓ 1202ms | http |
| 121.43.196.213:8222 | ✓ 879ms | ✓ 1022ms | ✓ 902ms | ✓ 1093ms | ✓ 867ms | http |
| 114.55.226.123:10086 | ✓ 1045ms | ✓ 1393ms | ✓ 966ms | ✓ 1187ms | ✓ 1008ms | http |
| 209.126.84.232:8888 | ✓ 1064ms | ✓ 1555ms | ✓ 1500ms | ✓ 1500ms | ✓ 1066ms | http |
| 38.34.179.38:8449 | ✓ 745ms | ✓ 857ms | ✓ 259ms | ✓ 669ms | ✓ 939ms | http |
| 38.145.220.9:8448 | ✓ 712ms | ✓ 609ms | ✓ 238ms | ✓ 650ms | ✓ 1039ms | http |
| 38.145.203.98:8447 | ✓ 1337ms | 否 | ✓ 1907ms | ✓ 1714ms | 否 | http |
| 38.34.179.25:8453 | 否 | ✓ 1140ms | ✓ 1788ms | ✓ 1989ms | 否 | http |
| 38.145.208.151:8453 | ✓ 1196ms | ✓ 1588ms | ✓ 214ms | ✓ 787ms | ✓ 1736ms | http |
| 38.34.179.66:8446 | ✓ 1247ms | 否 | ✓ 712ms | ✓ 1394ms | ✓ 1528ms | http |
| 38.34.179.62:8450 | ✓ 1247ms | ✓ 1874ms | ✓ 822ms | 否 | ✓ 1020ms | http |
| 103.52.114.95:3128 | ✓ 822ms | 否 | ✓ 1078ms | ✓ 1298ms | ✓ 1074ms | http |
| 38.145.220.96:8445 | ✓ 894ms | ✓ 939ms | ✓ 97ms | ✓ 814ms | ✓ 1112ms | http |
| 38.145.203.98:8446 | ✓ 1156ms | ✓ 886ms | ✓ 270ms | ✓ 938ms | 否 | http |
| 38.145.203.105:8446 | ✓ 1196ms | ✓ 890ms | ✓ 257ms | ✓ 932ms | 否 | http |
| 38.145.203.107:8446 | ✓ 1127ms | ✓ 1001ms | ✓ 259ms | ✓ 994ms | 否 | http |
| 121.230.9.252:1080 | ✓ 1122ms | 否 | ✓ 1159ms | ✓ 1605ms | ✓ 1380ms | http |
| 195.123.209.48:3128 | ✓ 1742ms | 否 | ✓ 1560ms | 否 | ✓ 1849ms | http |
| 38.34.179.100:8452 | ✓ 908ms | 否 | ✓ 241ms | ✓ 987ms | ✓ 1913ms | http |
| 38.145.203.105:8447 | ✓ 1743ms | ✓ 750ms | ✓ 536ms | ✓ 1929ms | ✓ 1054ms | http |
| 45.136.130.197:8452 | ✓ 1356ms | ✓ 1915ms | ✓ 470ms | 否 | ✓ 1699ms | http |
| 16.78.119.130:443 | 否 | 否 | ✓ 1916ms | ✓ 1698ms | ✓ 1772ms | http |
| 38.34.183.47:8452 | ✓ 588ms | ✓ 1247ms | ✓ 930ms | ✓ 1764ms | ✓ 599ms | http |
| 45.136.198.40:3128 | ✓ 1282ms | ✓ 1723ms | 否 | 否 | ✓ 1907ms | http |
| 106.117.208.101:7890 | ✓ 1621ms | ✓ 1554ms | ✓ 1507ms | 否 | ✓ 1069ms | http |
| 38.34.179.155:8448 | ✓ 91ms | 否 | ✓ 817ms | ✓ 1430ms | ✓ 498ms | http |
| 160.238.65.6:3128 | ✓ 1746ms | ✓ 1511ms | ✓ 1229ms | 否 | 否 | http |
| 207.254.71.62:8088 | ✓ 1074ms | 否 | ✓ 898ms | 否 | ✓ 1563ms | http |
| 5.104.87.17:8051 | ✓ 1244ms | 否 | ✓ 883ms | ✓ 1354ms | ✓ 1344ms | http |
| 157.20.128.247:3125 | ✓ 1179ms | 否 | ✓ 1354ms | ✓ 1511ms | ✓ 1330ms | http |
| 116.80.49.166:3172 | ✓ 1836ms | 否 | 否 | ✓ 1805ms | ✓ 1624ms | http |

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
