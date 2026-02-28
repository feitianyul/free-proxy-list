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

最后更新：2026-02-28 14:40:29 UTC（2026-02-28 22:40:29 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 301ms | 否 | ✓ 1145ms | ✓ 1333ms | ✓ 1032ms | http |
| 3.213.157.4:3128 | ✓ 396ms | 否 | ✓ 1049ms | ✓ 1913ms | ✓ 1367ms | http |
| 94.177.131.12:3128 | ✓ 1450ms | 否 | ✓ 1395ms | ✓ 1109ms | ✓ 631ms | http |
| 34.78.200.22:3128 | ✓ 954ms | 否 | ✓ 1492ms | 否 | ✓ 1745ms | http |
| 5.101.0.233:3128 | ✓ 927ms | 否 | ✓ 1885ms | 否 | ✓ 1834ms | http |
| 132.145.93.138:1080 | ✓ 1490ms | 否 | ✓ 1793ms | ✓ 1339ms | 否 | http |
| 20.27.11.248:8561 | ✓ 523ms | ✓ 807ms | ✓ 451ms | ✓ 865ms | ✓ 663ms | http |
| 20.27.14.220:8561 | ✓ 564ms | ✓ 778ms | ✓ 456ms | ✓ 871ms | ✓ 682ms | http |
| 35.225.22.61:80 | ✓ 725ms | ✓ 1854ms | ✓ 400ms | ✓ 1149ms | 否 | http |
| 20.27.15.111:8561 | ✓ 494ms | ✓ 1783ms | ✓ 456ms | ✓ 789ms | ✓ 669ms | http |
| 20.78.118.91:8561 | ✓ 652ms | ✓ 854ms | ✓ 773ms | ✓ 847ms | ✓ 762ms | http |
| 120.92.212.16:7890 | ✓ 961ms | ✓ 1218ms | ✓ 1011ms | ✓ 1356ms | ✓ 971ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1455ms | ✓ 1149ms | ✓ 1614ms | 否 | http |
| 20.78.26.206:8561 | ✓ 1488ms | ✓ 1105ms | ✓ 606ms | 否 | 否 | http |
| 144.31.25.69:21064 | 否 | 否 | ✓ 589ms | ✓ 1931ms | ✓ 1364ms | http |
| 34.78.177.18:3128 | ✓ 1004ms | ✓ 1833ms | ✓ 916ms | ✓ 1893ms | ✓ 1608ms | http |
| 34.185.159.217:3128 | ✓ 804ms | 否 | ✓ 1002ms | 否 | ✓ 1748ms | http |
| 193.124.225.175:1080 | ✓ 752ms | 否 | ✓ 1446ms | 否 | ✓ 1804ms | http |
| 35.241.222.101:3128 | ✓ 1106ms | 否 | ✓ 1133ms | 否 | ✓ 1991ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1175ms | 否 | ✓ 1680ms | ✓ 944ms | http |
| 138.124.53.25:7443 | ✓ 918ms | 否 | ✓ 1624ms | 否 | ✓ 1760ms | http |
| 52.188.28.218:3128 | ✓ 1216ms | 否 | ✓ 1454ms | ✓ 1977ms | ✓ 1422ms | http |
| 20.210.39.153:8561 | ✓ 1725ms | 否 | ✓ 538ms | ✓ 847ms | ✓ 708ms | http |
| 36.147.78.166:80 | ✓ 1789ms | ✓ 1595ms | 否 | ✓ 1945ms | ✓ 1564ms | http |
| 162.240.154.26:3128 | ✓ 1024ms | ✓ 1747ms | ✓ 756ms | 否 | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1230ms | ✓ 1682ms | ✓ 1491ms | ✓ 1972ms | http |
| 210.223.44.230:3128 | ✓ 1220ms | ✓ 1154ms | 否 | ✓ 1134ms | ✓ 845ms | http |
| 195.123.209.48:3128 | ✓ 1351ms | 否 | ✓ 1515ms | 否 | ✓ 1767ms | http |
| 81.70.169.194:80 | ✓ 1934ms | 否 | ✓ 968ms | 否 | ✓ 1345ms | http |
| 45.140.147.82:1082 | ✓ 635ms | 否 | ✓ 625ms | ✓ 1644ms | ✓ 1231ms | http |
| 45.140.147.82:1081 | ✓ 643ms | ✓ 1396ms | ✓ 1236ms | 否 | ✓ 1018ms | http |
| 34.7.88.87:3128 | ✓ 868ms | 否 | ✓ 1184ms | 否 | ✓ 1812ms | http |
| 142.171.85.32:1080 | ✓ 800ms | ✓ 1562ms | 否 | ✓ 919ms | ✓ 1704ms | http |
| 121.230.9.26:1080 | ✓ 1761ms | ✓ 1691ms | ✓ 1211ms | 否 | 否 | http |
| 35.234.17.221:8080 | 否 | ✓ 1276ms | 否 | ✓ 972ms | ✓ 849ms | http |
| 34.158.73.60:3128 | ✓ 757ms | 否 | ✓ 1344ms | 否 | ✓ 1747ms | http |
| 121.237.181.137:8888 | 否 | ✓ 1151ms | ✓ 913ms | 否 | ✓ 916ms | http |
| 121.230.8.181:1080 | ✓ 1116ms | 否 | ✓ 1738ms | ✓ 1601ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1003ms | 否 | ✓ 1031ms | ✓ 1441ms | ✓ 1040ms | http |
| 121.43.196.210:8222 | ✓ 1135ms | ✓ 1030ms | ✓ 873ms | ✓ 1112ms | ✓ 853ms | http |
| 121.43.196.213:8222 | ✓ 1208ms | ✓ 1035ms | ✓ 830ms | ✓ 1074ms | ✓ 890ms | http |
| 114.55.226.123:10086 | 否 | ✓ 1454ms | ✓ 994ms | ✓ 1243ms | ✓ 1025ms | http |
| 168.235.110.63:3128 | ✓ 907ms | ✓ 1400ms | ✓ 1144ms | ✓ 1340ms | ✓ 999ms | http |
| 165.227.5.10:8888 | ✓ 517ms | 否 | 否 | ✓ 720ms | ✓ 649ms | http |
| 45.136.198.40:3128 | ✓ 864ms | 否 | ✓ 1690ms | 否 | ✓ 1700ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1055ms | ✓ 1982ms | ✓ 1334ms | http |
| 36.147.78.166:443 | ✓ 1775ms | 否 | ✓ 1647ms | ✓ 1903ms | 否 | http |
| 54.88.116.133:80 | ✓ 319ms | 否 | ✓ 853ms | ✓ 1589ms | ✓ 1358ms | http |
| 3.214.214.245:80 | ✓ 1831ms | 否 | ✓ 798ms | ✓ 1513ms | ✓ 1536ms | http |
| 14.56.107.244:3128 | ✓ 766ms | 否 | ✓ 1616ms | 否 | ✓ 755ms | http |
| 65.108.203.37:28080 | ✓ 1382ms | 否 | ✓ 1535ms | 否 | ✓ 1977ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1668ms | ✓ 1433ms | ✓ 1503ms | http |
| 100.50.110.44:80 | ✓ 1480ms | 否 | ✓ 544ms | ✓ 1484ms | 否 | http |
| 100.51.102.147:80 | ✓ 1477ms | 否 | ✓ 529ms | ✓ 1528ms | 否 | http |
| 217.76.245.80:999 | 否 | 否 | ✓ 1330ms | ✓ 1591ms | ✓ 1334ms | http |
| 91.236.238.103:8080 | ✓ 1683ms | 否 | 否 | ✓ 1799ms | ✓ 1561ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 905ms | ✓ 1654ms | ✓ 880ms | http |
| 59.46.216.131:30001 | ✓ 991ms | ✓ 1984ms | ✓ 1051ms | 否 | 否 | http |
| 121.40.231.103:7890 | ✓ 963ms | ✓ 1162ms | ✓ 1335ms | 否 | ✓ 1268ms | http |
| 223.16.170.103:3128 | ✓ 1097ms | 否 | ✓ 1280ms | 否 | ✓ 1108ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 750ms | ✓ 1129ms | ✓ 663ms | http |
| 172.212.68.37:3128 | ✓ 374ms | 否 | ✓ 1227ms | ✓ 1644ms | ✓ 1028ms | http |

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
