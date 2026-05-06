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

最后更新：2026-05-06 15:16:35 UTC（2026-05-06 23:16:35 UTC+8）

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
| 218.108.131.186:17890 | ✓ 1014ms | ✓ 1065ms | ✓ 804ms | ✓ 1053ms | ✓ 875ms | http |
| 113.160.132.26:8080 | ✓ 1564ms | 否 | ✓ 1584ms | ✓ 1300ms | ✓ 1408ms | http |
| 115.231.181.40:8128 | ✓ 965ms | 否 | ✓ 853ms | ✓ 1891ms | ✓ 1836ms | http |
| 120.92.108.86:7890 | ✓ 1234ms | 否 | ✓ 1687ms | ✓ 1896ms | 否 | http |
| 45.125.67.37:8443 | ✓ 849ms | 否 | ✓ 1185ms | ✓ 1197ms | ✓ 1021ms | http |
| 43.156.132.113:3128 | ✓ 1353ms | 否 | ✓ 998ms | ✓ 988ms | ✓ 844ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1630ms | ✓ 1019ms | ✓ 1165ms | 否 | http |
| 223.16.170.103:80 | ✓ 1482ms | 否 | ✓ 1314ms | ✓ 1084ms | ✓ 1098ms | http |
| 8.219.97.248:80 | ✓ 1100ms | 否 | ✓ 957ms | ✓ 1709ms | 否 | http |
| 138.197.68.35:4857 | 否 | ✓ 1235ms | 否 | ✓ 1487ms | ✓ 991ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1853ms | ✓ 1275ms | ✓ 1256ms | http |
| 47.77.216.82:1080 | ✓ 1112ms | ✓ 830ms | ✓ 190ms | ✓ 767ms | ✓ 695ms | http |
| 84.47.150.125:1080 | ✓ 1037ms | 否 | ✓ 1891ms | 否 | ✓ 1755ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1825ms | ✓ 1413ms | 否 | ✓ 1584ms | http |
| 94.131.118.129:1081 | ✓ 1193ms | ✓ 1742ms | ✓ 930ms | ✓ 1571ms | ✓ 1599ms | http |
| 34.96.238.40:8080 | ✓ 1536ms | ✓ 1597ms | 否 | ✓ 951ms | ✓ 975ms | http |
| 150.107.140.238:3128 | ✓ 1433ms | 否 | 否 | ✓ 1090ms | ✓ 1874ms | http |
| 147.45.186.28:3128 | ✓ 1468ms | 否 | ✓ 1980ms | 否 | ✓ 1672ms | http |
| 94.131.118.129:1082 | ✓ 1157ms | 否 | ✓ 652ms | ✓ 1649ms | ✓ 1062ms | http |
| 113.118.54.215:7890 | ✓ 870ms | ✓ 1685ms | 否 | ✓ 1925ms | ✓ 1734ms | http |
| 8.154.21.175:3128 | ✓ 843ms | ✓ 1002ms | ✓ 803ms | ✓ 1103ms | ✓ 908ms | http |
| 65.108.203.35:18080 | ✓ 910ms | ✓ 1974ms | ✓ 1677ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1210ms | ✓ 1423ms | ✓ 1456ms | 否 | 否 | http |
| 1.231.81.166:3128 | 否 | ✓ 1174ms | ✓ 1484ms | ✓ 1817ms | ✓ 892ms | http |
| 168.110.52.228:3128 | ✓ 466ms | 否 | ✓ 522ms | ✓ 733ms | ✓ 599ms | http |
| 116.118.48.147:3128 | 否 | 否 | ✓ 1017ms | ✓ 1161ms | ✓ 979ms | http |
| 154.90.48.209:9090 | ✓ 1612ms | 否 | 否 | ✓ 1330ms | ✓ 1866ms | http |
| 181.78.77.210:999 | ✓ 1464ms | 否 | ✓ 1558ms | ✓ 1982ms | 否 | http |
| 43.133.44.89:8888 | 否 | 否 | ✓ 1712ms | ✓ 1166ms | ✓ 1505ms | http |
| 62.113.119.14:8080 | ✓ 1624ms | 否 | ✓ 948ms | ✓ 1710ms | ✓ 1272ms | http |
| 101.6.65.112:10080 | ✓ 970ms | ✓ 1185ms | ✓ 996ms | ✓ 1174ms | ✓ 934ms | http |
| 3.101.133.120:80 | ✓ 268ms | ✓ 1285ms | ✓ 1138ms | ✓ 838ms | ✓ 1019ms | http |
| 194.59.247.34:10808 | ✓ 696ms | ✓ 1770ms | ✓ 1532ms | 否 | ✓ 1580ms | http |
| 47.112.25.109:7890 | 否 | ✓ 1575ms | 否 | ✓ 1759ms | ✓ 1641ms | http |
| 113.176.92.71:3128 | ✓ 1465ms | ✓ 1297ms | ✓ 1310ms | ✓ 1188ms | ✓ 913ms | http |
| 217.182.195.221:30003 | ✓ 1103ms | 否 | ✓ 914ms | 否 | ✓ 1774ms | http |
| 121.230.8.235:1080 | ✓ 1179ms | ✓ 1512ms | ✓ 1976ms | ✓ 1997ms | 否 | http |
| 61.52.131.172:8443 | ✓ 894ms | ✓ 1239ms | ✓ 1015ms | ✓ 1188ms | ✓ 945ms | http |
| 121.230.8.144:1080 | ✓ 1134ms | ✓ 1406ms | ✓ 1828ms | 否 | 否 | http |
| 103.172.70.173:8080 | ✓ 1850ms | 否 | ✓ 1508ms | ✓ 1433ms | ✓ 1253ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1019ms | ✓ 1069ms | ✓ 1205ms | http |

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
