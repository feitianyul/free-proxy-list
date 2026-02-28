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

最后更新：2026-02-28 17:27:19 UTC（2026-03-01 01:27:19 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1599ms | 否 | ✓ 989ms | ✓ 1352ms | 否 | http |
| 3.213.157.4:3128 | ✓ 613ms | ✓ 1872ms | ✓ 289ms | ✓ 1172ms | ✓ 1284ms | http |
| 168.235.110.63:3128 | ✓ 582ms | 否 | ✓ 1078ms | ✓ 1747ms | ✓ 1064ms | http |
| 195.123.209.48:3128 | ✓ 1138ms | 否 | ✓ 1547ms | 否 | ✓ 1731ms | http |
| 115.231.181.40:8128 | ✓ 901ms | 否 | ✓ 1778ms | ✓ 1849ms | ✓ 1819ms | http |
| 221.127.195.224:8888 | 否 | 否 | ✓ 1091ms | ✓ 1201ms | ✓ 1061ms | http |
| 103.84.95.54:7890 | ✓ 984ms | 否 | ✓ 641ms | ✓ 1076ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1154ms | ✓ 1478ms | ✓ 1197ms | ✓ 1408ms | ✓ 1889ms | http |
| 142.171.85.32:1080 | ✓ 737ms | 否 | ✓ 901ms | ✓ 899ms | 否 | http |
| 36.147.78.166:80 | ✓ 1610ms | 否 | ✓ 1684ms | ✓ 1791ms | 否 | http |
| 81.70.169.194:80 | ✓ 920ms | ✓ 1200ms | ✓ 1632ms | ✓ 1179ms | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1309ms | ✓ 1327ms | 否 | ✓ 995ms | http |
| 103.104.99.29:80 | ✓ 1978ms | 否 | ✓ 1552ms | ✓ 1741ms | ✓ 1574ms | http |
| 103.104.99.89:80 | ✓ 1977ms | 否 | ✓ 1447ms | ✓ 1559ms | ✓ 1932ms | http |
| 172.104.63.237:3128 | ✓ 1689ms | 否 | ✓ 1987ms | 否 | ✓ 1989ms | http |
| 120.92.212.16:7890 | ✓ 894ms | ✓ 1368ms | ✓ 932ms | ✓ 1405ms | ✓ 1154ms | http |
| 193.124.225.175:1080 | ✓ 789ms | 否 | ✓ 1196ms | 否 | ✓ 1661ms | http |
| 165.225.120.17:11702 | ✓ 881ms | 否 | ✓ 880ms | ✓ 1858ms | ✓ 1438ms | http |
| 165.225.120.17:12497 | ✓ 877ms | 否 | ✓ 872ms | ✓ 1847ms | ✓ 1497ms | http |
| 121.237.181.137:8888 | ✓ 1710ms | 否 | ✓ 910ms | ✓ 1215ms | 否 | http |
| 165.225.120.17:10906 | ✓ 893ms | 否 | ✓ 879ms | 否 | ✓ 1749ms | http |
| 165.225.120.17:11912 | ✓ 895ms | 否 | ✓ 891ms | ✓ 1833ms | ✓ 1483ms | http |
| 165.225.120.17:11745 | ✓ 886ms | 否 | ✓ 875ms | ✓ 1931ms | ✓ 1497ms | http |
| 165.225.120.17:10458 | ✓ 889ms | 否 | ✓ 870ms | ✓ 1843ms | ✓ 1741ms | http |
| 165.225.120.17:11995 | ✓ 869ms | 否 | ✓ 870ms | ✓ 1854ms | ✓ 1444ms | http |
| 103.82.23.118:5171 | ✓ 1912ms | 否 | ✓ 1315ms | ✓ 1930ms | ✓ 1932ms | http |
| 165.225.120.17:12215 | ✓ 875ms | 否 | ✓ 870ms | ✓ 1861ms | 否 | http |
| 165.225.120.17:10919 | ✓ 886ms | 否 | ✓ 889ms | ✓ 1836ms | 否 | http |
| 45.140.147.155:1081 | ✓ 671ms | 否 | ✓ 1387ms | 否 | ✓ 1310ms | http |
| 59.46.216.131:30001 | ✓ 936ms | 否 | 否 | ✓ 1746ms | ✓ 1070ms | http |
| 222.228.171.92:8080 | ✓ 1917ms | ✓ 1756ms | 否 | ✓ 980ms | ✓ 1209ms | http |
| 35.234.17.221:8080 | ✓ 1653ms | 否 | ✓ 1008ms | ✓ 986ms | 否 | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1649ms | ✓ 1533ms | ✓ 1908ms | http |
| 103.236.64.247:8888 | ✓ 1956ms | ✓ 1209ms | 否 | ✓ 1629ms | 否 | http |
| 165.225.120.17:10880 | ✓ 1881ms | 否 | ✓ 1914ms | 否 | ✓ 1770ms | http |
| 165.225.120.17:12585 | ✓ 928ms | 否 | ✓ 1068ms | ✓ 1834ms | ✓ 1402ms | http |
| 165.225.120.17:11070 | ✓ 928ms | 否 | ✓ 1109ms | ✓ 1835ms | ✓ 1424ms | http |
| 165.225.120.17:11178 | ✓ 927ms | 否 | ✓ 1077ms | ✓ 1864ms | ✓ 1430ms | http |
| 165.225.120.17:11099 | ✓ 926ms | 否 | ✓ 1036ms | 否 | ✓ 1412ms | http |
| 165.225.120.17:12265 | ✓ 928ms | 否 | ✓ 1108ms | ✓ 1907ms | ✓ 1652ms | http |
| 165.225.120.17:12693 | ✓ 1695ms | 否 | ✓ 1063ms | ✓ 1839ms | ✓ 1416ms | http |
| 35.225.22.61:80 | ✓ 1215ms | 否 | ✓ 560ms | ✓ 1132ms | ✓ 1177ms | http |
| 36.147.78.166:443 | ✓ 1765ms | ✓ 1594ms | 否 | ✓ 1898ms | 否 | http |
| 103.133.223.21:8080 | 否 | 否 | ✓ 1784ms | ✓ 1861ms | ✓ 1349ms | http |
| 1.12.62.237:8080 | ✓ 1650ms | ✓ 1994ms | ✓ 1851ms | 否 | 否 | http |
| 85.208.108.43:2094 | ✓ 580ms | 否 | ✓ 1467ms | ✓ 1302ms | ✓ 1150ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1149ms | ✓ 1162ms | ✓ 1329ms | ✓ 1012ms | http |
| 138.124.53.25:7443 | ✓ 1305ms | 否 | 否 | ✓ 1860ms | ✓ 1402ms | http |
| 148.135.85.87:1080 | ✓ 1267ms | ✓ 1826ms | ✓ 797ms | ✓ 869ms | ✓ 687ms | http |
| 165.225.120.17:10728 | ✓ 925ms | 否 | ✓ 1052ms | ✓ 1854ms | ✓ 1444ms | http |
| 70.61.188.34:3128 | ✓ 835ms | 否 | ✓ 1981ms | ✓ 1754ms | 否 | http |
| 45.140.147.82:1081 | ✓ 604ms | 否 | ✓ 1490ms | ✓ 1893ms | 否 | http |
| 45.140.147.82:1082 | ✓ 598ms | ✓ 1350ms | ✓ 1424ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1261ms | 否 | ✓ 1983ms | 否 | ✓ 1758ms | http |
| 210.77.18.31:7890 | ✓ 996ms | ✓ 992ms | ✓ 860ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 1048ms | ✓ 1040ms | ✓ 861ms | ✓ 1074ms | ✓ 850ms | http |
| 34.205.52.219:80 | ✓ 422ms | ✓ 1635ms | ✓ 1192ms | ✓ 1388ms | ✓ 1385ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1523ms | ✓ 1308ms | ✓ 1290ms | http |
| 123.20.24.166:8118 | ✓ 1747ms | 否 | 否 | ✓ 1748ms | ✓ 1034ms | http |
| 165.225.120.17:11828 | ✓ 1632ms | 否 | 否 | ✓ 1914ms | ✓ 1420ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1529ms | ✓ 1605ms | ✓ 1408ms | http |
| 34.179.157.99:3128 | ✓ 1133ms | 否 | ✓ 1576ms | 否 | ✓ 1911ms | http |
| 52.54.20.49:80 | 否 | ✓ 1624ms | ✓ 307ms | ✓ 1652ms | ✓ 1553ms | http |
| 54.88.116.133:80 | 否 | 否 | ✓ 1017ms | ✓ 1488ms | ✓ 928ms | http |

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
