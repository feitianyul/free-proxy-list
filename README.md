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

最后更新：2026-04-29 11:01:33 UTC（2026-04-29 19:01:33 UTC+8）

**代理总数：41**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 41 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 41 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 727ms | ✓ 980ms | ✓ 922ms | ✓ 869ms | ✓ 821ms | http |
| 34.71.229.255:3128 | ✓ 485ms | 否 | ✓ 1277ms | ✓ 1096ms | ✓ 1041ms | http |
| 218.108.131.186:17890 | ✓ 1021ms | ✓ 1311ms | 否 | ✓ 1322ms | ✓ 842ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1861ms | ✓ 930ms | ✓ 1489ms | ✓ 1057ms | http |
| 34.96.238.40:8080 | ✓ 1618ms | 否 | 否 | ✓ 973ms | ✓ 1480ms | http |
| 45.167.124.71:999 | ✓ 1415ms | 否 | ✓ 1177ms | ✓ 1886ms | ✓ 1471ms | http |
| 217.76.245.80:999 | ✓ 1143ms | 否 | ✓ 1422ms | ✓ 1927ms | ✓ 1655ms | http |
| 154.64.232.35:8080 | ✓ 928ms | 否 | ✓ 1944ms | 否 | ✓ 1346ms | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 946ms | ✓ 1204ms | ✓ 1914ms | http |
| 47.253.71.133:8081 | ✓ 786ms | 否 | ✓ 1493ms | 否 | ✓ 901ms | http |
| 80.92.204.47:1081 | ✓ 641ms | ✓ 1333ms | 否 | ✓ 1934ms | 否 | http |
| 163.227.146.38:8080 | ✓ 1796ms | 否 | ✓ 1607ms | ✓ 1861ms | 否 | http |
| 46.101.95.183:8888 | ✓ 1886ms | 否 | ✓ 947ms | ✓ 1980ms | 否 | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1314ms | ✓ 1941ms | ✓ 1540ms | http |
| 168.110.52.228:3128 | ✓ 524ms | 否 | ✓ 1275ms | ✓ 755ms | ✓ 596ms | http |
| 217.60.252.32:13057 | ✓ 1598ms | 否 | ✓ 1521ms | 否 | ✓ 1456ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1899ms | ✓ 1715ms | ✓ 1233ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1775ms | 否 | ✓ 1451ms | ✓ 1559ms | ✓ 1655ms | http |
| 77.110.119.136:3128 | 否 | 否 | ✓ 599ms | ✓ 1699ms | ✓ 1224ms | http |
| 183.238.3.150:7897 | ✓ 904ms | ✓ 1087ms | ✓ 977ms | ✓ 1114ms | ✓ 825ms | http |
| 210.223.44.230:3128 | ✓ 1509ms | ✓ 1879ms | ✓ 1097ms | ✓ 1988ms | ✓ 1453ms | http |
| 8.154.21.175:3128 | ✓ 920ms | ✓ 1170ms | ✓ 1818ms | ✓ 1064ms | ✓ 936ms | http |
| 121.230.8.41:1080 | ✓ 1256ms | ✓ 1855ms | ✓ 1968ms | ✓ 1632ms | ✓ 1107ms | http |
| 101.32.244.83:8080 | ✓ 1006ms | ✓ 1796ms | ✓ 928ms | ✓ 1354ms | ✓ 1278ms | http |
| 121.43.196.210:8222 | ✓ 936ms | ✓ 1053ms | ✓ 814ms | ✓ 1122ms | ✓ 942ms | http |
| 121.43.196.213:8222 | ✓ 912ms | ✓ 1071ms | ✓ 850ms | ✓ 1154ms | ✓ 930ms | http |
| 114.55.226.123:10086 | ✓ 1123ms | ✓ 1463ms | ✓ 1058ms | ✓ 1340ms | ✓ 1163ms | http |
| 62.113.119.14:8080 | ✓ 800ms | 否 | ✓ 918ms | ✓ 1706ms | ✓ 1262ms | http |
| 45.59.122.132:80 | ✓ 1239ms | 否 | ✓ 808ms | 否 | ✓ 1700ms | http |
| 47.101.159.19:8899 | ✓ 904ms | ✓ 996ms | ✓ 969ms | ✓ 1135ms | ✓ 897ms | http |
| 159.89.31.62:8080 | ✓ 652ms | 否 | ✓ 1695ms | 否 | ✓ 1533ms | http |
| 172.236.145.31:7890 | ✓ 1551ms | 否 | 否 | ✓ 1617ms | ✓ 1906ms | http |
| 103.170.196.74:8080 | 否 | 否 | ✓ 1520ms | ✓ 1556ms | ✓ 1252ms | http |
| 183.232.248.73:7890 | ✓ 1699ms | 否 | ✓ 1295ms | ✓ 1798ms | 否 | http |
| 77.110.116.224:3128 | ✓ 1241ms | 否 | ✓ 1266ms | 否 | ✓ 1792ms | http |
| 152.32.132.190:7890 | ✓ 1251ms | ✓ 1569ms | ✓ 1593ms | 否 | ✓ 853ms | http |
| 8.219.97.248:80 | ✓ 1800ms | 否 | ✓ 1412ms | ✓ 1322ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1198ms | ✓ 1737ms | ✓ 1516ms | ✓ 1654ms | ✓ 1272ms | http |
| 45.153.231.229:8080 | ✓ 1925ms | 否 | ✓ 1958ms | 否 | ✓ 1994ms | http |
| 61.52.131.172:8443 | ✓ 934ms | ✓ 1265ms | ✓ 1029ms | ✓ 1199ms | ✓ 919ms | http |
| 103.39.51.207:8080 | ✓ 1546ms | 否 | 否 | ✓ 1882ms | ✓ 1775ms | http |

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
